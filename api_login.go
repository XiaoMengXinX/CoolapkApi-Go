package coolapk

import (
	"context"
	"fmt"
	ent "github.com/XiaoMengXinX/CoolapkApi-Go/entities"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type LoginCaptcha struct {
	Image []byte
	ID    string
}

func (c *Coolapk) LoginByPasswordWithCtx(user, password, captcha, captchaID string, ctx context.Context) (*ent.LoginData, *LoginCaptcha, error) {
	var result ent.LoginData

	url := "https://account.coolapk.com/auth/loginByCoolApk"
	header, body, err := c.loginRequest("GET", url, "", "", ctx)
	if err != nil {
		return &result, nil, err
	}

	var sessID string
	cookies := header.Values("set-cookie")
	for _, s := range cookies {
		if isMatch, _ := regexp.MatchString("SESSID", s); isMatch {
			sessID = strings.ReplaceAll(regexp.MustCompile("SESSID=((.*)+[0-9])").FindString(s), "SESSID=", "")
			break
		}
	}
	if sessID == "" {
		return &result, nil, fmt.Errorf("Get SESSID failed ")
	}
	c.Cookie = fmt.Sprintf("SESSID=%s; forward=https%%3A%%2F%%2Fwww.coolapk.com", sessID)

	if captcha != "" && captchaID != "" {
		c.Cookie = fmt.Sprintf("%s; captcha=%s", c.Cookie, captchaID)
	}

	rand.Seed(time.Now().UnixNano())
	var randomNum string
	for i := 0; i < 17; i++ {
		randomNum += strconv.Itoa(rand.Intn(10))
	}
	m := regexp.MustCompile(`requestHash : '(.*)',`)
	requestHash := m.FindString(string(body))
	if requestHash == "" {
		return &result, nil, fmt.Errorf("Get requestHash failed ")
	}
	requestHash = requestHash[15:29]

	c.Client = &loginClient{}
	err = c.Client.Request(c, &result, "POST", url, "", ctx, map[string]interface{}{
		"submit":       1,
		"login":        user,
		"password":     password,
		"randomNumber": "0undefined" + randomNum,
		"requestHash":  requestHash,
		"captcha":      captcha,
		"code":         "",
	})
	c.Client = &CoolapkClient{}

	if result.Status == -1 {
		var captchaData LoginCaptcha
		header, captchaData.Image, err = c.loginRequest("GET", "https://account.coolapk.com/auth/showCaptchaImage?"+strconv.FormatInt(time.Now().Unix(), 10), "", "", ctx)
		if err != nil {
			return &result, nil, err
		}
		for _, s := range header.Values("Set-Cookie") {
			ar := strings.Split(strings.Split(s, ";")[0], "=")
			if ar[0] == "captcha" {
				captchaData.ID = ar[1]
				break
			}
		}
		result.CaptchaID = captchaData.ID
		return &result, &captchaData, nil
	}
	return &result, nil, err
}

func (c *Coolapk) LoginByPassword(user, password, captcha, captchaID string) (*ent.LoginData, *LoginCaptcha, error) {
	return c.LoginByPasswordWithCtx(user, password, captcha, captchaID, context.Background())
}

type loginClient struct{}

func (d *loginClient) Request(c *Coolapk, result APIResp, method, url, _ string, ctx context.Context, paramters map[string]interface{}) error {
	body := parseParamters(paramters)
	header, resp, err := c.loginRequest(method, url, "", body, ctx)
	result.Deserialize(header, string(resp))
	cookies := map[string]string{}
	for _, s := range header.Values("Set-Cookie") {
		ar := strings.Split(strings.Split(s, ";")[0], "=")
		cookies[ar[0]] = ar[1]
	}
	if cookies["uid"] == "" || cookies["username"] == "" || cookies["token"] == "" {
		return fmt.Errorf("Get cookies failed ")
	}
	c.Cookie = fmt.Sprintf("uid=%s; username=%s; token=%s", cookies["uid"], cookies["username"], cookies["token"])
	return err
}

func (c *Coolapk) loginRequest(method, url, _, body string, ctx context.Context) (header http.Header, response []byte, err error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, strings.NewReader(body))
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 10; SM-G981B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.162 Mobile Safari/537.36")
	req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if method == "POST" {
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
	}
	req.Header.Set("Cookie", c.Cookie)

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	return resp.Header, respBody, err
}
