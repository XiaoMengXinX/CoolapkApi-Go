package feed

import (
	"fmt"
	coolapk "github.com/XiaoMengXinX/CoolapkApi-Go"
	"net/http"
	"strconv"
	"strings"
)

func FeedDetail(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(getArg(r, "id"))

	api := coolapk.New()
	api.Cookie = r.Header.Get("Cookie")

	result, err := api.GetFeedDetail(id)
	if err != nil {
		w.WriteHeader(500)
	}
	w = writeHeader(result.Header, w, r)

	_, _ = fmt.Fprintf(w, result.Response)
}

func getArg(r *http.Request, name string) string {
	var arg string
	values := r.URL.Query()
	arg = values.Get(name)
	return arg
}

func writeHeader(h http.Header, w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	for s, a := range h {
		for _, i := range a {
			if s == "Set-Cookie" {
				w.Header().Add(s, strings.ReplaceAll(i, "coolapk.com", r.Host))
				continue
			}
			w.Header().Add(s, i)
		}
	}
	return w
}
