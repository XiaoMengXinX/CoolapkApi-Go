package ent

type LoginData struct {
	RawData
	ErrorMsg
	CaptchaID      string `json:"captcha_id"`
	CaptchaURL     string `json:"captcha_url"`
	ThemeName      string `json:"themeName"`
	Charset        string `json:"charset"`
	RequestApp     string `json:"requestApp"`
	BaseUrl        string `json:"baseUrl"`
	RequestBase    string `json:"requestBase"`
	RequestBaseUrl string `json:"requestBaseUrl"`
	RequestHash    string `json:"requestHash"`
	RequestTime    int    `json:"requestTime"`
	AjaxRequest    bool   `json:"ajaxRequest"`
	RequestAjax    bool   `json:"requestAjax"`
	SESSION        struct {
		Uid                  int         `json:"uid"`
		Username             interface{} `json:"username"`
		Status               int         `json:"status"`
		IsLogin              bool        `json:"isLogin"`
		IsMember             bool        `json:"isMember"`
		IsAdministrator      bool        `json:"isAdministrator"`
		IsSuperAdministrator bool        `json:"isSuperAdministrator"`
		IsSubAdmin           int         `json:"isSubAdmin"`
		IsTester             bool        `json:"isTester"`
	} `json:"SESSION"`
	AppName          string `json:"appName"`
	ClassName        string `json:"className"`
	ControllerName   string `json:"controllerName"`
	CalledMethodName string `json:"calledMethodName"`
	Forward          string `json:"forward"`
	DisplayMode      string `json:"displayMode"`
	IosLogin         bool   `json:"iosLogin"`
	IsCommunity      bool   `json:"isCommunity"`
	CoolMarketClient bool   `json:"coolMarketClient"`
	Status           int    `json:"status"`
	MessageStatus    int    `json:"messageStatus"`
	Error            string `json:"error"`
	Message          string `json:"message"`
}
