package api

import (
	"fmt"
	"github.com/psanford/memfs"
	"io/fs"
	"log"
	"net/http"
)

var f *memfs.FS

func Captcha(w http.ResponseWriter, r *http.Request) {
	captchaID := GetArg(r, "id")
	file, err := fs.ReadFile(f, fmt.Sprintf("captcha/%s.jpg", captchaID))
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
