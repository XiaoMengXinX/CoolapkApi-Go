package api

import (
	"encoding/json"
	"fmt"
	coolapk "github.com/XiaoMengXinX/CoolapkApi-Go"
	"github.com/psanford/memfs"
	"io/fs"
	"log"
	"net/http"
	"os"
)

var FS = memfs.New()
var ocrAPI = os.Getenv("OCR_API")

func Login(w http.ResponseWriter, r *http.Request) {
	user := GetArg(r, "user")
	password := GetArg(r, "password")
	captcha := GetArg(r, "captcha")
	captchaID := GetArg(r, "captchaID")

	if captchaID != "" && captcha == "" && user == "" && password == "" {
		file, err := fs.ReadFile(FS, fmt.Sprintf("captcha/%s.jpg", captchaID))
		if err != nil {
			log.Println(err)
			WriteError(w, http.StatusInternalServerError, fmt.Errorf("captcha not found"))
			return
		}

		if err != nil {
			WriteError(w, http.StatusInternalServerError, fmt.Errorf("captcha not found"))
			return
		}
		w.Header().Set("Content-Type", "image/jpeg")

		_, _ = w.Write(file)
		return
	}

	if user == "" && password == "" {
		WriteError(w, http.StatusBadRequest, fmt.Errorf("invaid user or password"))
		return
	}

	c := coolapk.New()

	if captchaID != "" {
		cookie, _ := fs.ReadFile(FS, fmt.Sprintf("captcha/%s.txt", captchaID))
		c.Cookie = string(cookie)
	}

	result, captchaData, err := c.LoginByPassword(user, password, captcha, captchaID)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Add("Content-type", "application/json; charset=utf-8")

	if captchaData != nil {
		ocrResult, _ := UploadFile(ocrAPI, captchaData.Image)
		result, captchaData, err = c.LoginByPassword(user, password, string(ocrResult), captchaID)
		if err != nil {
			WriteError(w, http.StatusInternalServerError, err)
			return
		}
	}

	if captchaData != nil {
		_ = FS.MkdirAll("captcha", 0777)
		_ = FS.WriteFile(fmt.Sprintf("captcha/%s.jpg", captchaData.ID), captchaData.Image, 0755)
		_ = FS.WriteFile(fmt.Sprintf("captcha/%s.txt", captchaData.ID), []byte(c.Cookie), 0755)

		if r.TLS != nil {
			result.CaptchaURL = fmt.Sprintf("https://%s/login?captchaID=%s", r.Host, captchaData.ID)
		}
		result.CaptchaURL = fmt.Sprintf("http://%s/login?captchaID=%s", r.Host, captchaData.ID)
		resp, _ := json.Marshal(result)
		_, _ = fmt.Fprint(w, string(resp))
		return
	}

	w = WriteHeader(result.Header, w, r)

	_, _ = fmt.Fprint(w, result.Response)
}
