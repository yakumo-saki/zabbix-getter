package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Result  []hostResult
	Error   string
	Id      int
}

type hostResult struct {
	Hostid string
	Host   string
}

// Authenticate to zabbix and get authenticate token
func GetHostId(url string, token string, hostname string) (string, error) {
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
				"name": "%s"
			}
		},
		"id": 2,
		"auth": "%s"
	}`
	jsonStr := fmt.Sprintf(jsonTemplate, hostname, token)
	logger.T("Response\n", jsonStr)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", &ZabbixError{Msg: "Error while API request. (host.get)", Err: err}
	}

	// expected result
	// {"jsonrpc":"2.0","result":"057466f9a6cb65b3d57d9460cc792b9b","id":1}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	logger.T("Response\n", string(byteArray)) // htmlをstringで取得

	// parse JSON
	var decode_data hostApiResult
	if err := json.Unmarshal(byteArray, &decode_data); err != nil {
		return "", &ZabbixError{Msg: "Error while parsing json.\n" + string(byteArray), Err: err}
	}

	// TODO error when multiple host returned

	// 表示
	fmt.Println(decode_data.Result[0].Hostid)
	return decode_data.Result[0].Hostid, nil
}
