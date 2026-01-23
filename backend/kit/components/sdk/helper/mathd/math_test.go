package mathd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {

	assert.True(t, Min(1, 1) == 1)
	assert.True(t, Min(1, 2) == 1)
	assert.True(t, Min(1, 5) == 1)
}

func TestToFixed(t *testing.T) {

	fmt.Println(ToFixed4(0.342345))
	fmt.Println(ToFixed4(20000.0000000))
}

func TestCmn(t *testing.T) {

	fmt.Println(Cmn(8, 4))
}
