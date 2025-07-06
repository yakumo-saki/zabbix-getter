package zabbix

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/yakumo-saki/zabbix-getter/ylog"
)

//
// item.get API accepts hostname, then these methods are not used
//

// https://www.zabbix.com/documentation/2.2/manual/api
// request
// {
//     "jsonrpc": "2.0",
//     "method": "host.get",
//     "params": {
//         "output": [
//             "hostid",
//             "host"
//         ],
//         "selectInterfaces": [
//             "interfaceid",
//             "ip"
//         ]
//     },
//     "id": 2,
//     "auth": "0424bd59b807674191e7d77572075f33"
// }

// response
// {
//     "jsonrpc": "2.0",
//     "result": [
//         {
//             "hostid": "10084",
//             "host": "Zabbix server",
//             "interfaces": [
//                 {
//                     "interfaceid": "1",
//                     "ip": "127.0.0.1"
//                 }
//             ]
//         }
//     ],
//     "id": 2
// }

type hostApiResult struct {
	Jsonrpc string
	Result  []hostResult `json:",omitempty"`
	Error   string
	Id      int
}

type hostResult struct {
	Hostid string
	Host   string
}

// Get hostid
func (c *Client) GetHostId(hostname string) (string, error) {
	var logger = ylog.GetLogger()

	jsonTemplate := `
	{
		"jsonrpc": "2.0",
		"method": "host.get",
		"params": {
			"output": [
				"hostid",
				"host"
			]
			,"filter": {
				"host": "%s"
			}
		},
		"id": 2,
		"auth": "%s"
	}`
	jsonStr := fmt.Sprintf(jsonTemplate, hostname, c.Token)
	logger.T("Response\n", jsonStr)

	resp, err := c.PostApi(jsonStr)
	if err != nil {
		return "", &ZabbixError{Msg: "Error while API request. (host.get)", Err: err}
	}

	// expected result
	// {"jsonrpc":"2.0","result":[{"hostid":"10307","host":"envboy_livingroom"}],"id":2}
	byteArray, _ := io.ReadAll(resp.Body)
	logger.T("Response\n", string(byteArray)) // htmlをstringで取得

	resp.Body.Close()

	// parse JSON
	var decode_data hostApiResult
	if err := json.Unmarshal(byteArray, &decode_data); err != nil {
		logger.D(err)
		logger.E(string(byteArray))
		return "", &ZabbixError{Msg: "Error while parsing json.\n" + string(byteArray), Err: err}
	}

	// TODO error when multiple or no host returned
	if len(decode_data.Result) != 1 {
		return "", &ZabbixError{
			Msg: fmt.Sprintf("No hosts or multiple hosts found: %d", len(decode_data.Result)),
		}
	}

	// 表示
	logger.D(fmt.Sprintf("HostId is %s", decode_data.Result[0].Hostid))
	return decode_data.Result[0].Hostid, nil
}
