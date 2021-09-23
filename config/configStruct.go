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
	return "ENDPOINT=" + c.Url + " USERNAME=" + c.Username + " PASSWORD=<HIDDEN>" +
		" HOSTNAME=" + c.Hostname + " KEY=" + c.Key +
		" OUTPUT=" + c.Output + " LOGLEVEL=" + c.Loglevel
}
