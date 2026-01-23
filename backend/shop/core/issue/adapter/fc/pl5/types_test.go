package pl5

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	ta, _ := time.LoadLocation("Asia/Shanghai")

	fmt.Println(ta)
}

func TestF3dTicket_Prize(t *testing.T) {

	var cases = []struct {
		Case     *Ticket
		Target   *Ticket
		Expected int
	}{}

	for _, c := range cases {
		assert.Equal(t, c.Expected, len(c.Case.prize(c.Target)))
	}

}
