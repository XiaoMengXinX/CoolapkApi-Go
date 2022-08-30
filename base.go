package coolapk

import (
	"context"
	"net/http"

	token "github.com/XiaoMengXinX/FuckCoolapkTokenV2"
)

const defaultAPIEndpoint = "https://api2.coolapk.com/v6"

type APIResp interface {
	Deserialize(header http.Header, resp string)
}

type APIClient interface {
	Request(c *Coolapk, result APIResp, method, path, body string, ctx context.Context, params map[string]interface{}) error
}

type Coolapk struct {
	FakeClient  FakeClientInfo
	APIEndpoint string
	DeviceID    string
	Token       string
	UserAgent   string
	Cookie      string
	Client      APIClient
}

type FakeClientInfo struct {
	AndroidVer  string
	SDKVer      string
	Model       string
	BuildNumber string
	AppVersion  string
	AppCode     string
}

func (c *Coolapk) init() {
	clientInfo := getFakeClientInfo()
	c.APIEndpoint = defaultAPIEndpoint
	c.UserAgent = createUA(userAgentTmpl, clientInfo)
	c.FakeClient = clientInfo
	c.DeviceID, _ = token.GetToken()
	c.Client = &CoolapkClient{}
}

func New() Coolapk {
	var c Coolapk
	c.init()
	return c
}
