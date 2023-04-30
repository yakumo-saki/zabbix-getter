package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/yakumo-saki/zabbix-getter/ylog"
)

// request
// {
//   "jsonrpc": "2.0",
//   "method": "history.get",
//   "params": {
// 	    "output": "extend",
// 	    "history": 0,
// 	    "itemids": "23296",
// 	    "sortfield": "clock",
// 	    "sortorder": "DESC",
// 	    "limit": 10
//   },
// "auth": "038e1d7b1735c6a5436ee9eae095879e",
// "id": 1
// }

// response
// {
// 	"itemid": "23296",
// 	"clock": "1351090996",
// 	"value": "0.085",
// 	"ns": "563157632"
// },

type GetHistoryApiResult struct {
	Jsonrpc string
	Result  []HistoryResult `json:",omitempty"`
	Error   string
	Id      int
}

type HistoryResult struct {
	Itemid string `json:"itemid"`
	Clock  string `json:"clock"`
	Value  string `json:"value"`
	Ns     string `json:"ns"`
}

// Authenticate to zabbix and get authenticate token
func GetHistory(url string, token string, itemId string) (HistoryResult, error) {
	var logger = ylog.GetLogger()

	jsonTemplate := `{
		"jsonrpc": "2.0",
		"method": "history.get",
		"params": {
			"output": "extend",
			"history": 0,
			"itemids": "%s",
			"sortfield": "clock",
			"sortorder": "DESC",
			"limit": 1
		},
		"auth": "%s",
		"id": 3
	}`
	jsonStr := fmt.Sprintf(jsonTemplate, itemId, token)

	logger.T("Request\n", jsonStr)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return HistoryResult{}, &ZabbixError{
			Msg: "Error while API request. (history.get)",
			Err: err,
		}
	}

	byteArray, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	logger.T("Response\n", string(byteArray)) // htmlをstringで取得

	// parse JSON
	var decode_data GetHistoryApiResult
	if err := json.Unmarshal(byteArray, &decode_data); err != nil {
		logger.D(err)
		logger.D(string(byteArray))
		return HistoryResult{}, &ZabbixError{Msg: "Error while parsing json.\n" + string(byteArray), Err: err}
	}

	// 表示
	if len(decode_data.Result) < 1 {
		return HistoryResult{}, &ZabbixError{
			Msg: "No historys found",
		}
	}

	return decode_data.Result[0], nil
}

// Get Latest value from history.
// returns latest value, clock(unixtime), error
func GetLatestHistoryValue(url string, token string, itemId string) (string, string, error) {
	history, err := GetHistory(url, token, itemId)
	if err != nil {
		return "", "", err
	}
	return history.Value, history.Clock, nil
}
