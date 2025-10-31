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
