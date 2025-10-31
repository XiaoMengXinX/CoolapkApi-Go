package coolapk

import (
	"context"
	"github.com/XiaoMengXinX/CoolapkApi-Go/token"
	"io"
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
	headers := token.GenerateHeaders()
	for k, v := range headers {
		req.Header.Set(k, v)
	}

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
	respBody, err := io.ReadAll(resp.Body)

	return resp.Header, respBody, err
}
