package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/yakumo-saki/zabbix-getter/YLogger"
	"github.com/yakumo-saki/zabbix-getter/global"
)

func setHelpMessage() {
	pflag.Usage = func() {
		o := os.Stderr
		fmt.Fprintf(o, "%s (%s)\n", os.Args[0], global.Url)
		fmt.Fprintf(o, "Version: %s\n", global.Version)
		fmt.Fprintf(o, "\n")
		fmt.Fprintf(o, "Usage:\n")
		pflag.PrintDefaults()
		fmt.Fprintf(o, "\n")
	}
}

func GetConfigFromCommandLine() *ConfigStruct {
	var cliOption ConfigStruct

	url := pflag.StringP("endpoint", "e", "", "Zabbix Server API endpoint url. example: http://192.168.0.20/api_jsonrpc.php")
	host := pflag.StringP("hostname", "s", "", "Zabbix Hostname")
	key := pflag.StringP("key", "k", "", "Zabbix Item Key")
	loglevel := pflag.StringP("loglevel", "l", "", "Loglevel TRACE<DEBUG<INFO<WARN<ERROR<FATAL")
	output := pflag.StringP("output", "o", "", "Output type [JSON | VALUE]")

	setHelpMessage()
	pflag.Parse()

	// todo get from .env -> ~/.config/zabbix-getter.conf -> cli option
	cliOption.Url = *url
	cliOption.Hostname = *host
	cliOption.Key = *key
	cliOption.Loglevel = *loglevel
	cliOption.Output = *output

	if cliOption.Loglevel != "" {
		YLogger.SetLogLevel(cliOption.Loglevel)
	}

	return &cliOption
}
