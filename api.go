package coolapk

import (
	"context"
	ent "github.com/XiaoMengXinX/CoolapkApi-Go/entities"
)

func (c *Coolapk) GetFeedDetailWithCtx(id int, ctx context.Context) (*ent.FeedDetail, error) {
	var result ent.FeedDetail
	err := c.Client.Request(c, &result, "GET", "/feed/detail", "", ctx,
		map[string]interface{}{
			"id": id,
		},
	)
	return &result, err
}

func (c *Coolapk) GetFeedDetail(id int) (*ent.FeedDetail, error) {
	return c.GetFeedDetailWithCtx(id, context.Background())
}

func (c *Coolapk) LikeFeedWithCtx(id int, ctx context.Context) (*ent.FeedLike, error) {
	var result ent.FeedLike
	err := c.Client.Request(c, &result, "POST", "/feed/like", "", ctx,
		map[string]interface{}{
			"id":     id,
			"detail": 0,
		},
	)
	return &result, err
}

func (c *Coolapk) LikeFeed(id int) (*ent.FeedLike, error) {
	return c.LikeFeedWithCtx(id, context.Background())
}
