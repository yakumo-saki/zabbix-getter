package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/yakumo-saki/zabbix-getter/ylog"
)

type Client struct {
	VersionString string // expected "7.0.0" after init()
	Url           string
	User          string
	Password      string

	//
	Token string
}

const NOT_INIT = "NOT_INITIALIZED"

func NewClient(url string, user string, password string) *Client {
	c := new(Client)
	c.Url = url
	c.User = user
	c.Password = password
	c.VersionString = NOT_INIT

	return c
}

func (c *Client) Init() {
	ver, err := c.GetZabbixVersion()
	if err != nil {
		panic("Could not get zabbix version")
	}

	c.VersionString = ver
}

// Logout from zabbix
func (c *Client) Logout() error {
	var logger = ylog.GetLogger()
	jsonTemplate := `{"jsonrpc": "2.0","method": "user.logout","params": [],"id": 9999,"auth":"%s"}`

	jsonStr := fmt.Sprintf(jsonTemplate, c.Token)
	logger.T("Sending", jsonStr)

	req, _ := http.NewRequest("POST", c.Url, bytes.NewBuffer([]byte(jsonStr)))
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

func (c *Client) IsBefore54() bool {
	if !c.IsInitialized() {
		panic("not intiialized!")
	}
	limitVer := "5.4.0"
	return limitVer < c.VersionString
}

func (c *Client) IsAfter70() bool {
	if !c.IsInitialized() {
		panic("not intiialized!")
	}
	limitVer := "7.0.0"
	return limitVer <= c.VersionString
}

func (c *Client) IsInitialized() bool {
	return c.VersionString == NOT_INIT
}
