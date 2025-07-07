package zabbix

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/yakumo-saki/zabbix-getter/ylog"
)

// {     "jsonrpc": "2.0",     "result": "7.0.0",     "id": 1 }
type apiInfoApiResult struct {
	Jsonrpc string
	Error   string
	Id      int
	Result  string `json:",omitempty"`
}

// Get zabbix version
func (c *Client) GetZabbixVersion() (string, *ZabbixError) {
	var logger = ylog.GetLogger()
	payload := `{"jsonrpc": "2.0","method": "apiinfo.version","params": [] ,"id": 0}`

	req, _ := http.NewRequest("POST", c.Url, bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic("zabbix response is not expected at (apiinfo.version)")
	}

	byteArray, _ := io.ReadAll(resp.Body)
	logger.T("Response\n", string(byteArray))
	resp.Body.Close()

	// parse JSON
	var decode_data apiInfoApiResult
	if err := json.Unmarshal(byteArray, &decode_data); err != nil {
		logger.D(err)
		logger.E(string(byteArray))
		return "", &ZabbixError{Msg: "Error while parsing json.\n" + string(byteArray), Err: err}
	}

	logger.D("Zabbix server version: " + decode_data.Result)
	return decode_data.Result, nil
}
