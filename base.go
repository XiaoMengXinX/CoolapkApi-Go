package coolapk

import (
	"context"
	"net/http"
)

const defaultAPIEndpoint = "https://api.coolapk.com/v6"

type APIResp interface {
	Deserialize(header http.Header, resp string, statusCode int)
}

type APIClient interface {
	Request(c *Coolapk, result APIResp, method, path, body string, ctx context.Context, params map[string]interface{}) error
}

type Coolapk struct {
	// deprecated
	FakeClient FakeClientInfo
	// deprecated
	DeviceID    string
	APIEndpoint string
	UserAgent   string
	Cookie      string
	Client      APIClient
}

type FakeClientInfo struct {
	ApiVersion  string
	SdkInt      string
	Model       string
	BuildNumber string
	AppVersion  string
	AppCode     string
}

func (c *Coolapk) init() {
	c.APIEndpoint = defaultAPIEndpoint
	c.UserAgent = randomUA()
	c.Client = &CoolapkClient{}
}

func New() Coolapk {
	var c Coolapk
	c.init()
	return c
}
