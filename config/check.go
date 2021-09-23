package config

import "errors"

// 設定のチェック、だめならerror。OKならnil
func CheckConfig(c *ConfigStruct) error {

	switch {
	case c.Username == "":
		return errors.New("please specify zabbix username")
	case c.Password == "":
		return errors.New("please specify zabbix password")
	case c.Url == "":
		return errors.New("please specify zabbix API endpoint")
	case c.Hostname == "":
		return errors.New("please specify zabbix hostname")
	case c.Key == "":
		return errors.New("please specify zabbix item key")
	case c.Loglevel == "":
		// OK
	}

	return nil
}
