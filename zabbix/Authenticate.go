package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/yakumo-saki/zabbix-getter/ylog"
)

type authenticateResult struct {
	Jsonrpc string
	Result  string
	Error   zabbixError
	Id      int
}

type logoutResult struct {
	Jsonrpc string
	Result  bool
	Id      int
}

type zabbixError struct {
	Code    int
	Message string
	Data    string
}

func Authenticate(url string, username string, password string) (string, error) {
	result, errcode, err := AuthenticateAfter54(url, username, password)
	if errcode == -32602 {
		// zabbix 5.4以前
		result, _, err = AuthenticateBefore54(url, username, password)
	}

	return result, err
}

// Authenticate to zabbix and get authenticate token
func AuthenticateAfter54(url string, username string, password string) (string, int, error) {
	var logger = ylog.GetLogger()

	// zabbix <= 5.2 username is user, zabbix 5.4 or newer, username params is username
	jsonTemplate := `{"jsonrpc":"2.0","method":"user.login","params":{"username":"%s","password":"%s"},"id":1}`
	jsonStr := fmt.Sprintf(jsonTemplate, username, password)
	logger.T("Sending", jsonStr) // htmlをstringで取得

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", -1, &ZabbixError{Msg: "Error while API request. (user.login)", Err: err}
	}

	// expected result
	// {"jsonrpc":"2.0","result":"057466f9a6cb65b3d57d9460cc792b9b","id":1}
	byteArray, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	logger.T("Response", string(byteArray)) // htmlをstringで取得

	// parse JSON
	var decode_data authenticateResult
	if err := json.Unmarshal(byteArray, &decode_data); err != nil {
		return "", -1, err
	}

	// check authorize success
	if decode_data.Error.Code != 0 {
		return "", -1, fmt.Errorf("login failed: error %d %s", decode_data.Error.Code, decode_data.Error.Data)
	}

	// 表示
	// logger.D(decode_data.Result)
	return decode_data.Result, 0, nil
}

// Authenticate to zabbix and get authenticate token
func AuthenticateBefore54(url string, username string, password string) (string, int, error) {
	var logger = ylog.GetLogger()

	// zabbix <= 5.2 username is user, zabbix 5.4 or newer, username params is username
	jsonTemplate := `{"jsonrpc":"2.0","method":"user.login","params":{"user":"%s","password":"%s"},"id":1, auth: null}`
	jsonStr := fmt.Sprintf(jsonTemplate, username, password)
	logger.T("Sending", jsonStr) // htmlをstringで取得

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", -1, &ZabbixError{Msg: "Error while API request. (user.login)", Err: err}
	}

	// expected result
	// {"jsonrpc":"2.0","result":"057466f9a6cb65b3d57d9460cc792b9b","id":1}
	byteArray, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	logger.T("Response", string(byteArray)) // htmlをstringで取得

	// parse JSON
	var decode_data authenticateResult
	if err := json.Unmarshal(byteArray, &decode_data); err != nil {
		return "", -1, err
	}

	// check authorize success
	if decode_data.Error.Code != 0 {
		return "", -1, fmt.Errorf("login failed: error %d %s", decode_data.Error.Code, decode_data.Error.Data)
	}

	// 表示
	// logger.D(decode_data.Result)
	return decode_data.Result, 0, nil
}

func Logout(url string, auth string) error {
	var logger = ylog.GetLogger()
	jsonTemplate := `{"jsonrpc": "2.0","method": "user.logout","params": [],"id": 9999,"auth":"%s"}`

	jsonStr := fmt.Sprintf(jsonTemplate, auth)
	logger.T("Sending", jsonStr)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return &ZabbixError{Msg: "Error while API request. (user.logout)", Err: err}
	}

	// expected result
	// {"jsonrpc": "2.0", "result": true, "id": 9999 }
	byteArray, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	logger.T("Response", string(byteArray))

	// parse JSON
	var decode_data logoutResult
	if err := json.Unmarshal(byteArray, &decode_data); err != nil {
		return err
	}

	if !decode_data.Result {
		return fmt.Errorf("zabbix response is false at (user.logout)")
	}
	return nil
}
