package adminrest

import (
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/core/db"
	"github.com/gin-gonic/gin"
	"strings"
)

func ListLocs(c *gin.Context) {

	locs := comp.SDK().Loc().Get()

	//for _, loc := range locs {
	//	//loc.Id = loc.Name
	//
	//	for _, child := range loc.Children {
	//
	//		child.Id = child.Name
	//
	//		for _, l := range child.Children {
	//			l.Id = l.Name
	//		}
	//	}
	//
	//}
	gind.OK(c, locs)
}

func ListBanks(c *gin.Context) {

	branchName := c.Query("branchName")
	branchNo := c.Query("branchNo")

	ctx := c.Request.Context()

	var bankId, branchId string
	if branchNo != "" {
		parts := strings.Split(branchNo, "-")
		if len(parts) != 2 {
			gind.BadRequestf(c, "invalid branchNo")
			return
		}

		bankId = parts[0]
		branchId = parts[1]
	}

	banks, _, err := new(db.Bank).List(ctx, db.ListBankParams{
		BankId: bankId, BranchNo: branchId,
		BranchNameLike: branchName,
	})
	if err != nil {
		slf.WithError(err).Errorw("List Banks err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, banks)
}
