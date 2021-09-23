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

	logger.T("Config = ", config)

	return config
}

func mergeConfigs(env *ConfigStruct, cli *ConfigStruct) *ConfigStruct {

	var conf ConfigStruct
	conf.Url = getOneValue(env.Url, cli.Url)
	conf.Username = getOneValue(env.Username, cli.Username)
	conf.Hostname = getOneValue(env.Hostname, cli.Hostname)
	conf.Password = getOneValue(env.Password, cli.Password)
	conf.Key = getOneValue(env.Key, cli.Key)
	conf.Output = getOneValue(env.Output, cli.Output)
	conf.Loglevel = getOneValue(env.Loglevel, cli.Loglevel)

	return &conf
}

func getOneValue(env string, cli string) string {
	switch {
	case env == "" && cli == "":
		return ""
	case cli != "":
		return cli
	case env != "" && cli == "":
		return env
	default:
		logger.F("Unknown condition ", env, cli)
		panic("Unknown condition")
	}
}

func getConfigFromCommandLine() *ConfigStruct {
	var cliOption ConfigStruct

	url := flag.String("e", "", "Zabbix Server API endpoint url. example: http://192.168.0.20/api_jsonrpc.php")
	host := flag.String("s", "", "Zabbix Hostname")
	key := flag.String("k", "", "Zabbix Item Key")
	loglevel := flag.String("loglevel", "", "Loglevel TRACE<DEBUG<INFO<WARN<ERROR<FATAL")
	output := flag.String("output", "VALUE", "Output type [VALUE | JSON]")

	flag.Parse()

	// todo get from .env -> ~/.config/zabbix-getter.conf -> cli option
	cliOption.Url = *url
	cliOption.Hostname = *host
	cliOption.Key = *key
	cliOption.Loglevel = *loglevel
	cliOption.Output = *output

	return &cliOption
}
