package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yakumo-saki/zabbix-getter/global"
	"github.com/yakumo-saki/zabbix-getter/ylog"
)

func showVersionMessage() {
	o := os.Stderr
	fmt.Fprintf(o, "%s (%s)\n", filepath.Base(os.Args[0]), global.Url)
	fmt.Fprintf(o, "Version: %s\n", global.Version)
}

func setHelpMessage() {
	pflag.Usage = func() {
		o := os.Stderr
		showVersionMessage()
		fmt.Fprintf(o, "\n")
		fmt.Fprintf(o, "Username and Password must be declared: \n")
		fmt.Fprintf(o, "* ~/.config/zabbix-getter.conf\n")
		fmt.Fprintf(o, "* "+filepath.Join(getExecuteDir(), CONFFILE)+"\n")
		fmt.Fprintf(o, "* environment value USERNAME and PASSWORD\n")
		fmt.Fprintf(o, "\n")
		fmt.Fprintf(o, "Usage:\n")
		pflag.PrintDefaults()
		fmt.Fprintf(o, "\n")
		fmt.Fprintf(o, "\n")
	}
}

func GetConfigFromCommandLine() *ConfigStruct {
	var cliOption ConfigStruct

	url := pflag.StringP("endpoint", "e", "", "Zabbix Server API endpoint url. example: http://192.168.0.20/api_jsonrpc.php")
	zbxServer := pflag.StringP("zabbix", "z", "", "Zabbix Server hostname or IP address. using http. example: 192.168.0.20")
	zbxHttpsServer := pflag.StringP("zabbix-https", "Z", "", "Zabbix Server hostname or IP address. using https. example: 192.168.0.20")
	host := pflag.StringP("hostname", "s", "", "Zabbix Hostname")
	key := pflag.StringP("key", "k", "", "Zabbix Item Key")
	loglevel := pflag.StringP("loglevel", "l", "", "Loglevel TRACE>DEBUG>INFO>WARN>ERROR>FATAL")
	output := pflag.StringP("output", "o", "", "Output type [JSON | VALUE]")
	version := pflag.BoolP("version", "v", false, "Show version")
	help := pflag.Bool("help", false, "Show this help")
	debug := pflag.Bool("debug", false, "debug mode. this has no function.")

	setHelpMessage()
	pflag.Parse()

	if *version {
		showVersionMessage()
		os.Exit(0)
	}

	if *help || len(os.Args) == 1 {
		pflag.Usage()
		os.Exit(0)
	}

	// todo get from .env -> ~/.config/zabbix-getter.conf -> cli option
	cliOption.Url = *url
	cliOption.Hostname = *host
	cliOption.Key = *key
	cliOption.Loglevel = *loglevel
	cliOption.Output = *output

	switch {
	case *zbxServer != "" && *zbxHttpsServer != "":
		panic("can not use both -z and -zs.")
	case *zbxServer != "":
		cliOption.Url = fmt.Sprintf("http://%s/api_jsonrpc.php", *zbxServer)
	case *zbxHttpsServer != "":
		cliOption.Url = fmt.Sprintf("https://%s/api_jsonrpc.php", *zbxHttpsServer)
	}

	if cliOption.Loglevel != "" {
		ylog.SetLogLevel(cliOption.Loglevel)
	}

	if *debug {
		ylog.GetLogger().D("debug mode running.")
	}

	return &cliOption
}
