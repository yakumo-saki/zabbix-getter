package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/yakumo-saki/zabbix-getter/YLogger"
	"github.com/yakumo-saki/zabbix-getter/config"
	"github.com/yakumo-saki/zabbix-getter/zabbix"
)

var logger = YLogger.GetLogger("main")
var Flags config.ConfigStruct // dotenv + flags

func main() {

	logger = YLogger.GetLogger("main")
	YLogger.SetLogLevel("WARN")
	YLogger.SetLogOutput("STDERR")

	cfg := config.LoadConfig()
	cfgerr := config.CheckConfig(cfg)
	if cfgerr != nil {
		logger.F(cfgerr)
		os.Exit(10)
		return
	}

	YLogger.SetLogLevel(cfg.Loglevel)

	token, autherr := zabbix.Authenticate(cfg.Url, cfg.Username, cfg.Password)
	if autherr != nil {
		logger.F(autherr)
		logger.F("Error occured at Authenticate")
		os.Exit(11)
		return
	}

	hostId, hosterr := zabbix.GetHostId(cfg.Url, token, cfg.Hostname)
	if hosterr != nil {
		logger.F(hosterr)
		logger.F("Error occured at GetHostId. Hostname is wrong ?")
		os.Exit(12)
		return
	}
	logger.D("Hostname is OK. ID=" + hostId)

	item, itemerr := zabbix.GetItem(cfg.Url, token, hostId, cfg.Key)
	if itemerr != nil {
		logger.F(itemerr)
		logger.F("Error occured at GetItemId")
		os.Exit(13)
		return
	}
	logger.D("Itemname is OK. ID=" + item.Itemid)

	// Overwrite item.lastvalue because zabbix sometime return lastValue = ""
	latestValue, latestClock, histerr := zabbix.GetLatestHistoryValue(cfg.Url, token, item.Itemid)
	if histerr != nil {
		logger.F(histerr)
		logger.F("Error occured at GetHistory")
		os.Exit(14)
		return
	}
	logger.D(latestValue)

	item.Lastvalue = latestValue
	item.Lastclock = latestClock

	if strings.ToUpper(cfg.Output) == "VALUE" {
		fmt.Println(string(item.Lastvalue))
	} else {
		json, err := json.Marshal(item)
		if err != nil {
			logger.F(err)
			panic(err)
		}
		fmt.Println(string(json))
	}
}
