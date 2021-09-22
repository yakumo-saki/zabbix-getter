package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"github.com/yakumo-saki/zabbix-getter/YLogger"
	"github.com/yakumo-saki/zabbix-getter/zabbix"
)

var logger YLogger.Logger
var Flags ConfigStruct // dotenv + flags

type ConfigStruct struct {
	Url      string
	Username string
	Password string
	Hostname string
	Key      string
	Json     string
	Loglevel string
}

func (c ConfigStruct) String() string {
	return "ZBX_URL=" + c.Url + " HOSTNAME=" + c.Hostname + " KEY=" + c.Key
}

func main() {

	logger = &YLogger.YLogger{}
	logger.SetLogLevel("WARN")

	config := loadConfig()
	cfgerr := checkConfig(config)
	if cfgerr != nil {
		logger.F(cfgerr)
		return
	}

	// todo get username / password from env, .env
	logger.D(config)

	token, autherr := zabbix.Authenticate(config.Url, "Admin", "zabbix")
	if autherr != nil {
		logger.F(autherr)
		logger.F("Error occured at Authenticate")
		return
	}

	item, itemerr := zabbix.GetItem(config.Url, token, config.Hostname, config.Key)
	if itemerr != nil {
		logger.F(itemerr)
		logger.F("Error occured at GetItemId")
		return
	}

	fmt.Println(item.Lastvalue)
}

func checkConfig(c *ConfigStruct) error {

	switch {
	case c.Url == "":
		return errors.New("Please specify zabbix API endpoint")
	case c.Hostname == "":
		return errors.New("-s option is not set. Please specify zabbix hostname")
	case c.Hostname == "":
		return errors.New("-s option is not set. Please specify zabbix hostname")
	}

	return nil
}

func loadConfig() *ConfigStruct {
	// var conf ConfigStruct
	env := getConfigFromDotEnv()
	cli := getConfigFromCommandLine()

	fmt.Println(env, cli)
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

func getConfigFromDotEnv() *ConfigStruct {
	var conf ConfigStruct
	var dotenv string

	confdir, configerr := os.UserConfigDir()
	if configerr != nil {
		logger.E(os.Stderr, configerr)
		return &conf
	}

	fmt.Println(filepath.Join(confdir, "zabbix-getter.conf"))

	enverr := godotenv.Load(dotenv)
	if enverr != nil {
		logger.F("Error loading .env file:", dotenv)
	}

	return &conf
}
