package ent

type FeedLike struct {
	RawData
	ErrorMsg
	Data struct {
		Count          int      `json:"count"`
		RecentLikeList []string `json:"recentLikeList"`
	} `json:"data"`
}
