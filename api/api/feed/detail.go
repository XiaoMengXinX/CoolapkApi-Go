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
		api.WriteError(w, err)
		return
	}
	w = api.WriteHeader(result.Header, w, r)
	w.Header().Add("Content-type", "application/json; charset=utf-8")
	_, _ = fmt.Fprint(w, result.Response)
}
