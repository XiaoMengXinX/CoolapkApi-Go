package feed

import (
	"api/api"
	"fmt"
	coolapk "github.com/XiaoMengXinX/CoolapkApi-Go"
	"net/http"
	"strconv"
)

func FeedDetail(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(api.GetArg(r, "id"))

	c := coolapk.New()
	c.Cookie = r.Header.Get("Cookie")

	result, err := c.GetFeedDetail(id)
	if err != nil {
		w.WriteHeader(500)
	}
	w = api.WriteHeader(result.Header, w, r)

	_, _ = fmt.Fprintf(w, result.Response)
}
