package set

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type SetTestStruct struct {
	k1 string
	k2 string
	k3 string
}

func (s *SetTestStruct) UniqueKey() string {
	return s.k1 + s.k2
}

func TestSet_Add(t *testing.T) {

	set := NewSet[*SetTestStruct]()

	assert.True(t, set.Add(&SetTestStruct{k1: "1", k2: "2"}))
	assert.False(t, set.Add(&SetTestStruct{k1: "1", k2: "2"}))
	assert.False(t, set.Add(&SetTestStruct{k1: "1", k2: "2", k3: "3"}))
	assert.True(t, set.Add(&SetTestStruct{k1: "1", k2: "22", k3: "3"}))

}

func TestSet_Iter(t *testing.T) {

	set := NewSet[*SetTestStruct]()

	set.Add(&SetTestStruct{k1: "1", k2: "2"})
	set.Add(&SetTestStruct{k1: "1", k2: "2"})
	set.Add(&SetTestStruct{k1: "1", k2: "2", k3: "3"})
	set.Add(&SetTestStruct{k1: "1", k2: "22", k3: "3"})

	for x := range set.Iter() {
		fmt.Println(x.UniqueKey())
	}

	slice := set.ToSlice()

	assert.Equal(t, 2, len(slice))

}
