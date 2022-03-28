package ent

type FeedLike struct {
	Response string `json:"-"`
	ErrorMsg
	Data struct {
		Count          int      `json:"count"`
		RecentLikeList []string `json:"recentLikeList"`
	} `json:"data"`
}
