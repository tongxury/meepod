package rest

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/gin-gonic/gin"
)

func ListMatchFilters(c *gin.Context) {

	category := c.Query("category")

	ctx := c.Request.Context()

	switch category {
	case enum.MatchCategory_Z14.Value:
		issue, err := new(service.IssueService).FindLatestIssueIndex(ctx, enum.ItemId_rx9)
		if err != nil {
			gind.Error(c, err)
			return
		}

		if issue == "" {
			gind.OK(c)
			return
		}

		var filters []gin.H

		for i := 0; i < 10; i++ {

			x := conv.String(conv.Int(issue) - i)

			filters = append(filters, gin.H{
				"name":  fmt.Sprintf("%s期", x),
				"value": x,
			})
		}

		gind.OK(c, filters)
		return
	}

	gind.OK(c)
}

func ListMatches(c *gin.Context) {

	category := c.Query("category")
	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	var matches types.Matches
	var total int64
	var err error

	switch category {
	case enum.MatchCategory_Z14.Value:
		issue := c.Query("issue")

		if issue == "" {
			index, err := new(service.IssueService).FindCurrentIssueIndex(ctx, enum.ItemId_rx9)
			if err != nil {
				gind.Error(c, err)
				return
			}

			issue = index
		}

		matches, total, err = new(service.MatchService).ListZ14Matches(ctx, issue, page, size)
	case enum.MatchCategory_Zjc.Value:
		matches, total, err = new(service.MatchService).ListZjcMatches(ctx, page, size)

	}

	if err != nil {
		slf.WithError(err).Errorw("ListMatches err", slf.String("category", category))
		gind.Error(c, err)
		return
	}

	gind.Page(c, matches, page, size, total)
}
