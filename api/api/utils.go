package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

type ErrorMsg struct {
	Error string `json:"error"`
}

func A(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, ``) }

func GetArg(r *http.Request, name string) string {
	var arg string
	if r.Method == "POST" {
		arg = r.FormValue(name)
	} else {
		arg = r.URL.Query().Get(name)
	}
	return arg
}

func WriteHeader(h http.Header, w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	for s, a := range h {
		for _, i := range a {
			if s == "Set-Cookie" || s == "set-cookie" {
				i = strings.ReplaceAll(i, "coolapk.com", r.Host)
			}
			w.Header().Add(s, i)
		}
	}
	return w
}

func WriteError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(ErrorMsg{Error: err.Error()})
}

func UploadFile(url string, file []byte) ([]byte, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", "captcha.jpg")
	if err != nil {
		return nil, err
	}
	_, err = fileWriter.Write(file)
	if err != nil {
		return nil, err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	return respBody, nil
}
