package main

import (
	"encoding/json"
	"fmt"
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
		return
	}

	YLogger.SetLogLevel(cfg.Loglevel)

	token, autherr := zabbix.Authenticate(cfg.Url, cfg.Username, cfg.Password)
	if autherr != nil {
		logger.F(autherr)
		logger.F("Error occured at Authenticate")
		return
	}

	item, itemerr := zabbix.GetItem(cfg.Url, token, cfg.Hostname, cfg.Key)
	if itemerr != nil {
		logger.F(itemerr)
		logger.F("Error occured at GetItemId")
		return
	}

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
