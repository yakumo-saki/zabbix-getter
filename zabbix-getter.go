package main

import (
	"flag"
	"fmt"

	"github.com/yakumo-saki/zabbix-getter/YLogger"
	"github.com/yakumo-saki/zabbix-getter/zabbix"
)

var logger YLogger.Logger
var zabbixItem zabbixItemStruct

type zabbixItemStruct struct {
	Url      string
	Username string
	Password string
	Hostname string
	Key      string
}

func (z zabbixItemStruct) String() string {
	return "ZBX_URL=" + z.Url + " HOSTNAME=" + z.Hostname + " KEY=" + z.Key
}

func main() {
	logger = &YLogger.YLogger{}
	logger.SetLogLevel("DEBUG")
	logger.I("Hello, Hello")

	url := flag.String("e", "", "Zabbix Server API endpoint url. example: http://192.168.0.20/api_jsonrpc.php")
	host := flag.String("s", "", "Zabbix Hostname")
	key := flag.String("k", "", "Zabbix Item Key")

	flag.Parse()

	// todo get from .env -> ~/.config/zabbix-getter.conf -> cli option
	zabbixItem.Url = *url
	zabbixItem.Hostname = *host
	zabbixItem.Key = *key

	switch {
	case zabbixItem.Url == "":
		logger.F("Please specify zabbix API endpoint")
		return
	case zabbixItem.Hostname == "":
		logger.F("-s option is not set. Please specify zabbix hostname")
		return
	case zabbixItem.Hostname == "":
		logger.F("-s option is not set. Please specify zabbix hostname")
		return
	}

	// todo get username / password from env, .env
	logger.D(zabbixItem)

	token, autherr := zabbix.Authenticate("http://10.1.0.10/api_jsonrpc.php", "Admin", "zabbix")
	if autherr != nil {
		logger.F(autherr)
		logger.F("Error occured at Authenticate")
		return
	}

	hostId, hosterr := zabbix.GetHostId("http://10.1.0.10/api_jsonrpc.php", token, zabbixItem.Hostname)
	if hosterr != nil {
		logger.F(hosterr)
		logger.F("Error occured at GetHostId")
		return
	}

	item, itemerr := zabbix.GetItem("http://10.1.0.10/api_jsonrpc.php", token, hostId, zabbixItem.Key)
	if itemerr != nil {
		logger.F(itemerr)
		logger.F("Error occured at GetItemId")
		return
	}

	fmt.Println(item.Lastvalue)
}
