package config

import "errors"

// 設定のチェック、だめならerror。OKならnil
func CheckConfig(c *ConfigStruct) error {

	// 未設定ならデフォルト値をセットするもの
	if c.Output == "" {
		c.Output = "JSON"
	}
	if c.Loglevel == "" {
		c.Output = "WARN"
	}

	if c.Username == "" {
		return errors.New("please specify zabbix username")
	}

	if c.Password == "" {
		return errors.New("please specify zabbix password")
	}

	if c.Url == "" {
		return errors.New("please specify zabbix API endpoint")
	}
	if c.Hostname == "" {
		return errors.New("please specify zabbix hostname")
	}
	if c.Key == "" {
		return errors.New("please specify zabbix item key")
	}

	return nil
}
