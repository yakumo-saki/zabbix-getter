package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func GetConfigFromDotEnv() *ConfigStruct {
	var conf ConfigStruct

	loadDotEnvFiles()

	conf.Hostname = os.Getenv("HOSTNAME")
	conf.Url = os.Getenv("URL")
	conf.Key = os.Getenv("KEY")
	conf.Output = os.Getenv("OUTPUT")
	conf.Loglevel = os.Getenv("LOGLEVEL")
	conf.Password = os.Getenv("PASSWORD")
	conf.Username = os.Getenv("USERNAME")

	logger.T("config from env:", conf)

	return &conf
}

func loadDotEnvFiles() error {
	var dotenv string
	CONFFILE := "zabbix-getter.conf"

	confdir, configerr := os.UserConfigDir()
	if configerr != nil {
		logger.E(os.Stderr, configerr)
		return configerr
	}

	dotenv = filepath.Join(confdir, CONFFILE)
	loadDotEnv(dotenv)

	dotenv = filepath.Join(getExecuteDir(), CONFFILE)
	loadDotEnv(dotenv)

	// show all environment
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		logger.T(pair[0], " => ", pair[1])
	}
	return nil
}

// 実行ファイルのあるディレクトリを取得
func getExecuteDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

// dotenvファイルをロードする
// 存在しなくてもスルー
func loadDotEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		logger.I("Error loading .env file:" + path)
	}
	logger.D(path + " loaded.")
}
