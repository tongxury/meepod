package eth

import (
	"fmt"
	"testing"
)

func TestEtherScanCrawlerClient_FindTokenMeta(t *testing.T) {

	c := NewEtherScanCrawlerClient()

	meta, _ := c.FindAddressMeta("0xdac17f958d2ee523a2206206994597c13d831ec7")

	fmt.Println(meta)
	meta1, _ := c.FindTxMeta("0x2f1c5c2b44f771e942a8506148e256f94f1a464babc938ae0690c6e34cd79190")

	fmt.Println(meta1)
}
