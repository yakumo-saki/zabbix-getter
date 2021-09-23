package config

// var logger = YLogger.GetLogger("zabbix")

// 設定をロードします
func LoadConfig() *ConfigStruct {
	// var conf ConfigStruct
	cli := GetConfigFromCommandLine()
	env := GetConfigFromDotEnv()

	config := mergeConfigs(env, cli)
	SetDefaultConfig(config)
	CheckConfig(config)

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
