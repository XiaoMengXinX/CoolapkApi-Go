package api

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

func Captcha(w http.ResponseWriter, r *http.Request) {
	captchaID := GetArg(r, "id")
	file, err := fs.ReadFile(FS, fmt.Sprintf("captcha/%s.jpg", captchaID))
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
}
