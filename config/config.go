package config

import (
	"flag"
)

// var logger = YLogger.GetLogger("zabbix")

// 設定をロードします
func LoadConfig() *ConfigStruct {
	// var conf ConfigStruct
	env := GetConfigFromDotEnv()
	cli := getConfigFromCommandLine()

	config := mergeConfigs(env, cli)

	return config
}

func mergeConfigs(env *ConfigStruct, cli *ConfigStruct) *ConfigStruct {
	return cli
}

func getConfigFromCommandLine() *ConfigStruct {
	var cliOption ConfigStruct

	url := flag.String("e", "", "Zabbix Server API endpoint url. example: http://192.168.0.20/api_jsonrpc.php")
	host := flag.String("s", "", "Zabbix Hostname")
	key := flag.String("k", "", "Zabbix Item Key")
	loglevel := flag.String("loglevel", "", "Loglevel TRACE<DEBUG<INFO<WARN<ERROR<FATAL")

	flag.Parse()

	// todo get from .env -> ~/.config/zabbix-getter.conf -> cli option
	cliOption.Url = *url
	cliOption.Hostname = *host
	cliOption.Key = *key
	cliOption.Loglevel = *loglevel

	return &cliOption
}
