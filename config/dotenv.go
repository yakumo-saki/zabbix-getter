package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/yakumo-saki/zabbix-getter/ylog"
)

const CONFFILE = "zabbix-getter.conf"

// ~/.config/zabbix-getter.conf
//
func LoadFromDotConfig() (*ConfigStruct, error) {
	var logger = ylog.GetLogger()

	confdir, configerr := os.UserConfigDir()
	if configerr != nil {
		var c ConfigStruct
		logger.E(os.Stderr, configerr)
		return &c, configerr
	}

	dotenv := filepath.Join(confdir, CONFFILE)
	cfg, _ := loadDotEnv(dotenv)

	return cfg, nil
}

// 実行時ファイルのディレクトリのある zabbix-getter.conf
//
func LoadFromExecDir() (*ConfigStruct, error) {
	dotenv := filepath.Join(getExecuteDir(), CONFFILE)
	cfg, _ := loadDotEnv(dotenv)

	return cfg, nil
}

// 実行時の環境変数からconfigを生成
//
func LoadFromEnvValue() *ConfigStruct {

	var conf ConfigStruct

	conf.Hostname = os.Getenv("HOSTNAME")
	conf.Url = os.Getenv("ENDPOINT")
	conf.Key = os.Getenv("KEY")
	conf.Output = os.Getenv("OUTPUT")
	conf.Loglevel = os.Getenv("LOGLEVEL")
	conf.Password = os.Getenv("PASSWORD")
	conf.Username = os.Getenv("USERNAME")

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
func loadDotEnv(path string) (*ConfigStruct, error) {
	var logger = ylog.GetLogger()

	var conf ConfigStruct

	m, err := godotenv.Read(path)
	if err != nil {
		logger.D("Error loading .env file" + path)
		return &conf, err
	}

	conf.Hostname = m["HOSTNAME"]
	conf.Url = m["ENDPOINT"]
	conf.Key = m["KEY"]
	conf.Output = m["OUTPUT"]
	conf.Loglevel = m["LOGLEVEL"]
	conf.Password = m["PASSWORD"]
	conf.Username = m["USERNAME"]

	logger.D(path + " loaded.")
	return &conf, nil
}
