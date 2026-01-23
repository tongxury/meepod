package ethutils

import (
	"fmt"
	"testing"
)

func TestGetFuncSelector(t *testing.T) {

	v := GetFuncSelector("balanceOf(address)")
	fmt.Println(v)
}

//0x70a082310000000000000000000000008894e0a0c962cb723c1976a4421c95949be2d4e3
//0xe2e4263a0000000000000000000000009696f59e4d72e237be84ffd425dcad154bf96976
