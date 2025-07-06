package zabbix

import (
	"bytes"
	"net/http"
)

func (c *Client) PostApi(jsonStr string) (*http.Response, error) {

	req, _ := http.NewRequest("POST", c.Url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	if c.IsAfter70() {
		req.Header.Set("Authori", "application/json")
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	return resp, err
}
