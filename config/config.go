package config

import "github.com/yakumo-saki/zabbix-getter/ylog"

// 設定をロードします
func LoadConfig() *ConfigStruct {
	var logger = ylog.GetLogger()

	// var conf ConfigStruct
	cli := GetConfigFromCommandLine()
	dotConfig, _ := LoadFromDotConfig() // Linux/macOS
	execConfig, _ := LoadFromExecDir()  // windows対策
	envConfig := LoadFromEnvValue()

	config := mergeConfigs(execConfig, dotConfig)
	config = mergeConfigs(config, envConfig)
	config = mergeConfigs(config, cli)

	SetDefaultConfig(config)

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
	var logger = ylog.GetLogger()

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
