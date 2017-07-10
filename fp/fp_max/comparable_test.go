package fp_max

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
)

func Test_comparable(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field int
	}
	should.Equal(TestObjectComparable{0, 3}, plz.Max(
		TestObjectComparable{0, 2},
		TestObjectComparable{0, 3},
		TestObjectComparable{0, 1}))
}

type TestObjectComparable struct {
	Field1 int
	Field2 int
}

func (obj TestObjectComparable) Compare(that interface{}) int {
	thisVal := obj.Field2
	thatVal := that.(TestObjectComparable).Field2
	if thisVal < thatVal {
		return -1
	} else if thisVal == thatVal {
		return 0
	} else {
		return 1
	}
}
