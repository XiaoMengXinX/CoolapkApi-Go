package coolapk

import (
	"context"
	"net/http"

	token "github.com/XiaoMengXinX/FuckCoolapkTokenV2"
)

const defaultAPIEndpoint = "https://api.coolapk.com/v6"

type APIResp interface {
	Deserialize(header http.Header, resp string)
}

type APIClient interface {
	Request(c *Coolapk, result APIResp, method, path, body string, ctx context.Context, params map[string]interface{}) error
}

type Coolapk struct {
	APIEndpoint string
	DeviceID    string
	UserAgent   string
	Cookie      string
	Client      APIClient
}

func (c *Coolapk) init() {
	c.APIEndpoint = defaultAPIEndpoint
	c.UserAgent = getRandomUA(userAgentTmpl)
	c.DeviceID, _ = token.GetToken()
	c.Client = &CoolapkClient{}
}

func New() Coolapk {
	var c Coolapk
	c.init()
	return c
}
