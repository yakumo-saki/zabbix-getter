package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/yakumo-saki/zabbix-getter/global"
	"github.com/yakumo-saki/zabbix-getter/ylog"
)

func showVersionMessage() {
	o := os.Stderr
	fmt.Fprintf(o, "%s (%s)\n", os.Args[0], global.Url)
	fmt.Fprintf(o, "Version: %s\n", global.Version)
}

func setHelpMessage() {
	pflag.Usage = func() {
		o := os.Stderr
		showVersionMessage()
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
	loglevel := pflag.StringP("loglevel", "l", "", "Loglevel TRACE>DEBUG>INFO>WARN>ERROR>FATAL")
	output := pflag.StringP("output", "o", "", "Output type [JSON | VALUE]")
	version := pflag.BoolP("version", "v", false, "Show version")

	setHelpMessage()
	pflag.Parse()

	if *version {
		showVersionMessage()
		os.Exit(0)
	}

	// todo get from .env -> ~/.config/zabbix-getter.conf -> cli option
	cliOption.Url = *url
	cliOption.Hostname = *host
	cliOption.Key = *key
	cliOption.Loglevel = *loglevel
	cliOption.Output = *output

	if cliOption.Loglevel != "" {
		ylog.SetLogLevel(cliOption.Loglevel)
	}

	return &cliOption
}
