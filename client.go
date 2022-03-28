package coolapk

import (
	"context"
	"net/url"
	"strconv"
)

type CoolapkClient struct{}

func (d *CoolapkClient) Request(c *Coolapk, result APIResp, method, path, body string, ctx context.Context, paramters map[string]interface{}) error {
	params := url.Values{}
	for key, value := range paramters {
		switch value.(type) {
		case string:
			params.Add(key, value.(string))
		case int:
			params.Add(key, strconv.Itoa(value.(int)))
		case int64:
			params.Add(key, strconv.Itoa(int(value.(int64))))
		}
	}
	data := params.Encode()
	header, resp, err := c.Request(method, path, data, body, ctx)
	result.Deserialize(header, string(resp))
	return err
}
