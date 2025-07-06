package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/imdario/mergo"
	"github.com/yakumo-saki/zabbix-getter/config"
	"github.com/yakumo-saki/zabbix-getter/ylog"
	"github.com/yakumo-saki/zabbix-getter/zabbix"
)

var logger = ylog.GetLogger()
var Flags config.ConfigStruct // dotenv + flags

func main() {

	ylog.Init()
	logger = ylog.GetLogger()
	ylog.SetLogLevel("WARN")
	ylog.SetLogOutput("STDERR")

	cfg := config.LoadConfig()
	cfgerr := config.CheckConfig(cfg)
	if cfgerr != nil {
		fmt.Fprint(os.Stderr, cfgerr.Error()+". see help for detail.\n")
		os.Exit(10)
		return
	}

	err := ylog.SetLogLevel(cfg.Loglevel)
	if err != nil {
		ylog.SetLogLevel("WARN")
	}

	zabbix := zabbix.NewClient(cfg.Url, cfg.Username, cfg.Password)

	zabbix.Init()

	autherr := zabbix.Authenticate()
	if autherr != nil {
		logger.F(autherr)
		logger.F("Error occured at Authenticate")
		os.Exit(11)
		return
	}

	hostId, hosterr := zabbix.GetHostId(cfg.Hostname)
	if hosterr != nil {
		logger.F(hosterr)
		logger.F("Error occured at GetHostId. Hostname is wrong ?")
		os.Exit(12)
		return
	}
	logger.D("Hostname is OK. ID=" + hostId)

	item, itemerr := zabbix.GetItem(hostId, cfg.Key)
	if itemerr != nil {
		logger.F(itemerr)
		logger.F("Error occured at GetItemId")
		os.Exit(13)
		return
	}
	logger.D("Itemname is OK. ID=" + item.Itemid)

	latestValue, latestClock := item.Lastvalue, item.Lastclock

	if latestClock == "0" {
		// Overwrite item.lastvalue because zabbix sometime return lastValue = ""
		logger.D("Lastclock on item is 0.(maybe last data is little old) Try getting history.")

		val, clk, histerr := zabbix.GetLatestHistoryValue(item.Itemid)
		if histerr != nil {
			logger.F(histerr)
			logger.F("Error occured at GetHistory")
			os.Exit(14)
			return
		}

		logger.D("Getting history success.")
		latestValue = val
		latestClock = clk
	}

	item.Lastvalue = latestValue
	item.Lastclock = latestClock

	// TODO result check
	zabbix.Logout()

	// 出力
	if strings.ToUpper(cfg.Output) == "VALUE" {
		fmt.Println(string(item.Lastvalue))
	} else {
		unixtime, e := strconv.ParseInt(item.Lastclock, 0, 0)
		timeStr := ""
		if e == nil {
			timeStr = fmt.Sprint(time.Unix(unixtime, 0))
		}

		r := make(map[string]interface{})
		mergo.Map(&r, item, mergo.WithOverride)
		r["lastclockString"] = timeStr

		json, err := json.Marshal(r)
		if err != nil {
			logger.F(err)
			panic(err)
		}
		fmt.Println(string(json))
	}

}
