package config

type ConfigStruct struct {
	Url      string
	Username string
	Password string
	Hostname string
	Key      string
	Output   string
	Loglevel string
}

func (c ConfigStruct) String() string {
	return "ZBX_URL=" + c.Url + " HOSTNAME=" + c.Hostname + " KEY=" + c.Key
}
