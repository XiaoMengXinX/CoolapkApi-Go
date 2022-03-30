package api

import (
	"encoding/json"
	"fmt"
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
