package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// request
// {
//     "jsonrpc": "2.0",
//     "method": "item.get",
//     "params": {
//         "output": "extend",
//         "hostids": "10084",
//         "search": {
//             "key_": "system"
//         },
//         "sortfield": "name"
//     },
//     "auth": "038e1d7b1735c6a5436ee9eae095879e",
//     "id": 1
// }

// response
// {
//     "jsonrpc": "2.0",
//     "result": [
//         {
//             "itemid": "23298",
//             "type": "0",
//             "snmp_community": "",
//             "snmp_oid": "",
//             "hostid": "10084",
//             "name": "Context switches per second",
//             "key_": "system.cpu.switches",
//             "delay": "60",
//             "history": "7",
//             "trends": "365",
//             "lastvalue": "2552",
//             "lastclock": "1351090998",
//             "prevvalue": "2641",
//             "state": "0",
//             "status": "0",
//             "value_type": "3",
//             "trapper_hosts": "",
//             "units": "sps",
//             "multiplier": "0",
//             "delta": "1",
//             "snmpv3_securityname": "",
//             "snmpv3_securitylevel": "0",
//             "snmpv3_authpassphrase": "",
//             "snmpv3_privpassphrase": "",
//             "formula": "1",
//             "error": "",
//             "lastlogsize": "0",
//             "logtimefmt": "",
//             "templateid": "22680",
//             "valuemapid": "0",
//             "delay_flex": "",
//             "params": "",
//             "ipmi_sensor": "",
//             "data_type": "0",
//             "authtype": "0",
//             "username": "",
//             "password": "",
//             "publickey": "",
//             "privatekey": "",
//             "mtime": "0",
//             "lastns": "564054253",
//             "flags": "0",
//             "filter": "",
//             "interfaceid": "1",
//             "port": "",
//             "description": "",
//             "inventory_link": "0",
//             "lifetime": "0"
//         },...
// 	]
// }

type GetItemApiResult struct {
	Jsonrpc string
	Result  []ItemResult
	Error   string
	Id      int
}

type ItemResult struct {
	Itemid    string
	Hostid    string
	Key_      string
	Name      string
	Lastvalue string
	Lastclock string
	Units     string
}

// Authenticate to zabbix and get authenticate token
func GetItem(url string, token string, hostname string, itemname string) (ItemResult, error) {
	jsonTemplate := `
	{
		"jsonrpc": "2.0",
		"method": "item.get",
		"params": {
			"output": [
				"itemid",
				"hostid",
				"key_",
				"lastvalue",
				"lastclock",
				"units"
			],
			"filter": {
				"key_": "%s",
				"host": "%s"
			}
		},
		"id": 3,
		"auth": "%s"
	}`
	jsonStr := fmt.Sprintf(jsonTemplate, itemname, hostname, token)

	fmt.Println(jsonStr)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return ItemResult{}, &ZabbixError{Msg: "Error while API request. (item.get)", Err: err}
	}

	// expected result
	// {"jsonrpc":"2.0","result":"057466f9a6cb65b3d57d9460cc792b9b","id":1}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(byteArray)) // htmlをstringで取得

	// parse JSON
	var decode_data GetItemApiResult
	if err := json.Unmarshal(byteArray, &decode_data); err != nil {
		fmt.Println(err)
		return ItemResult{}, &ZabbixError{Msg: "Error while parsing json.\n" + string(byteArray), Err: err}
	}

	// 表示
	fmt.Println(decode_data.Result)

	// check return item key is actualy match. (api is like search)

	return decode_data.Result[0], nil
}
