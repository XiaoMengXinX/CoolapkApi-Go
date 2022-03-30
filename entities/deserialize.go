package ent

import (
	"encoding/json"
	"net/http"
)

func (r *FeedDetail) Deserialize(header http.Header, resp string) {
	_ = json.Unmarshal([]byte(resp), r)
	r.Response = resp
	r.Header = header
}

func (r *FeedLike) Deserialize(header http.Header, resp string) {
	_ = json.Unmarshal([]byte(resp), r)
	r.Response = resp
	r.Header = header
}

func (r *LoginData) Deserialize(header http.Header, resp string) {
	_ = json.Unmarshal([]byte(resp), r)
	r.Response = resp
	r.Header = header
}
