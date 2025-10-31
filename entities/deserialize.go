package ent

import (
	"encoding/json"
	"net/http"
)

func (r *FeedDetail) Deserialize(header http.Header, resp string, statusCode int) {
	_ = json.Unmarshal([]byte(resp), r)
	r.Response = resp
	r.Header = header
	r.StatusCode = statusCode
}
