package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Captcha(w http.ResponseWriter, r *http.Request) {
	captchaID := GetArg(r, "id")
	file, err := FS.Open("captcha/" + captchaID)
	if err != nil {
		WriteError(w, fmt.Errorf("captcha not found"))
		return
	}
	defer file.Close()

	imgBytes, err := ioutil.ReadAll(file)
	if err != nil {
		WriteError(w, fmt.Errorf("captcha not found"))
		return
	}
	w.Header().Set("Content-Type", "image/jpeg")

	_, _ = w.Write(imgBytes)
}
