package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestN(t *testing.T) {

	ctx := context.Background()

	//https://docs.bscscan.com/misc-tools-and-utilities/public-rpc-nodes
	//c, _ := NewClient("https://cloudflare-eth.com")
	c, _ := NewClient("https://bsc-dataseed2.defibit.io/")
	tx, _, _ := c.TransactionByHash(ctx, common.HexToHash("0x36580dafc90c9409ea445762edae8c1a4f89ba232a8421121ecfd108a51df982"))
	//tx.To().String()

	//c, _ := NewEthClients([]string{"https://bsc-dataseed2.defibit.io/"})

	fmt.Println(tx.To().String())
}
