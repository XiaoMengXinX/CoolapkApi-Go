package ent

import (
	"encoding/json"
)

func (r *FeedDetail) Deserialize(resp string) {
	_ = json.Unmarshal([]byte(resp), r)
	r.Response = resp
}
