package coolapk

import (
	"context"
	token "github.com/XiaoMengXinX/FuckCoolapkTokenV2"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func (c *Coolapk) Request(method, path, param, body string, ctx context.Context) (response []byte, err error) {
	isPost := method == "POST"
	client := &http.Client{}

	var req *http.Request
	if isPost {
		req, err = http.NewRequestWithContext(ctx, method, c.APIEndpoint+path+"?"+param, strings.NewReader(body))
	} else {
		req, err = http.NewRequestWithContext(ctx, method, c.APIEndpoint+path+"?"+param, nil)
	}
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("X-App-Device", c.DeviceID)
	req.Header.Set("X-App-Token", token.GetTokenWithDeviceCode(c.DeviceID))
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("X-Sdk-Int", "31")
	req.Header.Set("X-Sdk-Locale", "zh-CN")
	req.Header.Set("X-App-Id", "com.coolapk.market")
	req.Header.Set("X-App-Version", "12.1")
	req.Header.Set("X-App-Code", "2203161")
	req.Header.Set("X-Api-Version", "12")
	req.Header.Set("X-App-Channel", "coolapk")
	req.Header.Set("X-App-Mode", "universal")
	req.Header.Set("X-App-Supported", "2203161")
	if isPost {
		req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	}
	if c.Cookie != "" {
		req.Header.Set("Cookie", c.Cookie)
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
