package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func GetConfigFromDotEnv() *ConfigStruct {
	var conf ConfigStruct
	var dotenv string
	CONFFILE := "zabbix-getter.conf"

	confdir, configerr := os.UserConfigDir()
	if configerr != nil {
		logger.E(os.Stderr, configerr)
		return &conf
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

	return &conf
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
