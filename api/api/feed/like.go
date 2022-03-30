package feed

import (
	"api/api"
	"fmt"
	coolapk "github.com/XiaoMengXinX/CoolapkApi-Go"
	"net/http"
	"strconv"
)

func FeedLike(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(api.GetArg(r, "id"))

	c := coolapk.New()
	c.Cookie = r.Header.Get("Cookie")

	result, err := c.LikeFeed(id)
	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w = api.WriteHeader(result.Header, w, r)

	_, _ = fmt.Fprint(w, result.Response)
}
