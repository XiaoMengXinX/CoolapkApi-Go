package coolapk

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func (c *Coolapk) Request(method, path, param, body string, ctx context.Context) (header http.Header, response []byte, err error) {
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
	req.Header.Set("X-App-Token", c.Token)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("X-Sdk-Int", c.FakeClient.SDKVer)
	req.Header.Set("X-Sdk-Locale", "zh-CN")
	req.Header.Set("X-App-Id", "com.coolapk.market")
	req.Header.Set("X-App-Version", c.FakeClient.AppVersion)
	req.Header.Set("X-App-Code", c.FakeClient.AppCode)
	req.Header.Set("X-Api-Version", c.FakeClient.AndroidVer)
	req.Header.Set("X-App-Channel", "coolapk")
	req.Header.Set("X-App-Mode", "universal")
	req.Header.Set("X-App-Supported", c.FakeClient.AppCode)
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
	respBody, err := ioutil.ReadAll(resp.Body)
	return resp.Header, respBody, err
}
