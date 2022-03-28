package feed

import (
	"fmt"
	coolapk "github.com/XiaoMengXinX/CoolapkApi-Go"
	"net/http"
	"strconv"
)

func FeedLike(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(getArg(r, "id"))

	api := coolapk.New()
	api.Cookie = r.Header.Get("Cookie")

	result, err := api.LikeFeed(id)
	if err != nil {
		w.WriteHeader(500)
	}
	w = writeHeader(result.Header, w, r)

	_, _ = fmt.Fprintf(w, result.Response)
}
