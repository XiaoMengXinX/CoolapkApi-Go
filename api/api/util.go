package api

import (
	"fmt"
	"net/http"
	"strings"
)

func A(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "") }

func GetArg(r *http.Request, name string) string {
	var arg string
	values := r.URL.Query()
	arg = values.Get(name)
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
