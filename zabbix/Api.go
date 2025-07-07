package zabbix

import (
	"bytes"
	"net/http"
	"regexp"
	"strings"

	"github.com/yakumo-saki/zabbix-getter/ylog"
)

func (c *Client) PostApi(jsonStr string) (*http.Response, error) {
	var logger = ylog.GetLogger()

	// replace auth on payload when zbx >= 7.0
	payload := jsonStr
	if c.IsAfter70() {
		payload = RemoveAuth(payload)
	}

	req, _ := http.NewRequest("POST", c.Url, bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	if c.IsAfter70() && c.Token != NOT_INIT {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	{
		p := jsonStr
		p = strings.ReplaceAll(p, c.Password, "<PASSWORD>")
		logger.T("Request", p) // htmlをstringで取得
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	return resp, err
}

func RemoveAuth(jsonStr string) string {
	var logger = ylog.GetLogger()
	logger.T(jsonStr)

	re := regexp.MustCompile(`"auth":\s?".*",?`)
	logger.T("Removing: '" + re.String() + "'")

	payload := jsonStr
	payload = re.ReplaceAllString(payload, "")
	// payload = strings.ReplaceAll(payload, auth+",", "")
	// payload = strings.ReplaceAll(payload, auth, "")

	logger.T(payload)

	return payload
}
