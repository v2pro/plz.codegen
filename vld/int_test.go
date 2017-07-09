package vld

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
)

func Test_int_required(t *testing.T) {
	t.Skip("WIP")
	should := require.New(t)

	type TestObject struct {
		MyField int `validate:"required"`
	}
	err := plz.Validate(TestObject{})
	should.NotNil(err)
	should.Contains(err.Error(), "MyField")

	err = plz.Validate(TestObject{1})
	should.Nil(err)
}
