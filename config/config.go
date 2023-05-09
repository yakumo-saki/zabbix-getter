package config

import (
	"runtime"
	"strings"

	"github.com/yakumo-saki/zabbix-getter/ylog"
)

// 設定をロードします
func LoadConfig() *ConfigStruct {
	var logger = ylog.GetLogger()

	var config *ConfigStruct

	cli := GetConfigFromCommandLine()

	if strings.ToUpper(runtime.GOOS) == "DARWIN" {
		dotConfig, _ := LoadFromXDGConfigHomeDir()
		config = dotConfig
	}

	dotConfig, _ := LoadFromDotConfig() // Linux: .config macOS:~/Library/Application Support

	execConfig, _ := LoadFromExecDir() // windows対策
	envConfig := LoadFromEnvValue()

	config = mergeConfigs(config, dotConfig)
	config = mergeConfigs(config, execConfig)
	config = mergeConfigs(config, envConfig)
	config = mergeConfigs(config, cli)

	SetDefaultConfig(config)

	logger.T("Config = ", config)

	return config
}

// baseの設定をoverwriteの設定で上書きする。
// ただし、overwrite側の設定が空文字列ではない場合のみ上書きする。
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

// 空文字列ではない値を取得する。
// base, overwrite双方とも空文字列ではない場合は overwrite を返す
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
