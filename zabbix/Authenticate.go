package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yakumo-saki/zabbix-getter/YLogger"
)

type authenticateResult struct {
	Jsonrpc string
	Result  string
	Id      int
}

var logger = YLogger.GetLogger("zabbix")

// Authenticate to zabbix and get authenticate token
func Authenticate(url string, username string, password string) (string, error) {
	jsonTemplate := `
	{"jsonrpc":"2.0","method":"user.login","params":
	{"user":"%s","password":"%s"},
	"id":1,"auth":null}`
	jsonStr := fmt.Sprintf(jsonTemplate, username, password)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", &ZabbixError{Msg: "Error while API request. (user.login)", Err: err}
	}

	// expected result
	// {"jsonrpc":"2.0","result":"057466f9a6cb65b3d57d9460cc792b9b","id":1}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	logger.D(string(byteArray)) // htmlをstringで取得

	// parse JSON
	var decode_data authenticateResult
	if err := json.Unmarshal(byteArray, &decode_data); err != nil {
		return "", err
	}

	// 表示
	logger.D(decode_data.Result)
	return decode_data.Result, nil
}
