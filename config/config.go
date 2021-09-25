package config

// var logger = YLogger.GetLogger("zabbix")

// 設定をロードします
func LoadConfig() *ConfigStruct {
	// var conf ConfigStruct
	cli := GetConfigFromCommandLine()
	dotConfig, _ := LoadFromDotConfig()
	execConfig, _ := LoadFromExecDir()
	envConfig := LoadFromEnvValue()

	config := mergeConfigs(execConfig, dotConfig)
	config = mergeConfigs(config, envConfig)
	config = mergeConfigs(config, cli)

	SetDefaultConfig(config)
	CheckConfig(config)

	logger.T("Config = ", config)

	return config
}

func mergeConfigs(base *ConfigStruct, overwrite *ConfigStruct) *ConfigStruct {

	var conf ConfigStruct
	conf.Url = getOneValue(base.Url, overwrite.Url)
	conf.Username = getOneValue(base.Username, overwrite.Username)
	conf.Hostname = getOneValue(base.Hostname, overwrite.Hostname)
	conf.Password = getOneValue(base.Password, overwrite.Password)
	conf.Key = getOneValue(base.Key, overwrite.Key)
	conf.Output = getOneValue(base.Output, overwrite.Output)
	conf.Loglevel = getOneValue(base.Loglevel, overwrite.Loglevel)

	return &conf
}

func getOneValue(base string, overwrite string) string {
	switch {
	case base == "" && overwrite == "":
		return ""
	case overwrite != "":
		return overwrite
	case base != "" && overwrite == "":
		return base
	default:
		logger.F("Unknown condition ", base, overwrite)
		panic("Unknown condition")
	}
}
