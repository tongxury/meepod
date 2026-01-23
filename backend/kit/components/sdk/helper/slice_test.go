package helper

import (
	"fmt"
	"testing"
)

func TestSplitSlice(t *testing.T) {

	src := []int{1, 2, 3, 4, 4, 5, 6, 8, 9, 9, 4}

	r := SplitSlice(src, 2)

	fmt.Println(r)
}
