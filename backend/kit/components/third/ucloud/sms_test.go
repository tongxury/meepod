package ucloud

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestSMSClient_SendSMS(t *testing.T) {
	var code = "2322"
	c, _ := NewSMSClient("7PT9vS5XSDpFa0n7MqlIsmINWsP50cTF21CrRJzDm5", "BiFJKxnx8raF7U7WZtZAm87PoKpEjG7PhApHGa8nTnZ749PRYYnMwHHC0tAYJLw6u5")

	ctx := context.Background()
	send, err := c.Send(ctx, "13391620292", code, "福彩科技", "UTA230808KQJPZH")
	if err != nil {
		log.Print(err)
	}

	fmt.Println(send)
}
