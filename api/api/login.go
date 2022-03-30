package api

import (
	"encoding/json"
	"fmt"
	coolapk "github.com/XiaoMengXinX/CoolapkApi-Go"
	"github.com/psanford/memfs"
	"io/fs"
	"log"
	"net/http"
)

var FS = memfs.New()

func Login(w http.ResponseWriter, r *http.Request) {
	id := GetArg(r, "id")
	user := GetArg(r, "user")
	password := GetArg(r, "password")
	captcha := GetArg(r, "captcha")
	captchaID := GetArg(r, "captchaID")

	if id != "" {
		file, err := fs.ReadFile(FS, fmt.Sprintf("captcha/%s.jpg", id))
		if err != nil {
			log.Println(err)
			WriteError(w, fmt.Errorf("captcha not found"))
			return
		}

		if err != nil {
			WriteError(w, fmt.Errorf("captcha not found"))
			return
		}
		w.Header().Set("Content-Type", "image/jpeg")

		_, _ = w.Write(file)
		return
	}

	c := coolapk.New()
	result, captchaData, err := c.LoginByPassword(user, password, captcha, captchaID)
	if err != nil {
		WriteError(w, err)
		return
	}
	w.Header().Add("Content-type", "application/json; charset=utf-8")

	if captchaData != nil {
		_ = FS.MkdirAll("captcha", 0777)
		err := FS.WriteFile(fmt.Sprintf("captcha/%s.jpg", captchaData.ID), captchaData.Image, 0755)
		if err != nil {
			result.Error = err.Error()
		}

		if r.TLS != nil {
			result.CaptchaURL = fmt.Sprintf("https://%s/captcha?id=%s", r.Host, captchaData.ID)
		}
		result.CaptchaURL = fmt.Sprintf("http://%s/captcha?id=%s", r.Host, captchaData.ID)
		resp, _ := json.Marshal(result)
		_, _ = fmt.Fprint(w, string(resp))
		return
	}

	w = WriteHeader(result.Header, w, r)

	_, _ = fmt.Fprint(w, result.Response)
}
