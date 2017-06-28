package wombat

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz/tags"
)

func Test_int_required(t *testing.T) {
	type TestObject struct {
		myField int
	}
	tags.Define(func(o *TestObject) tags.Tags {
		return tags.D(tags.S(),
			tags.F(&o.myField, "validate", "required"),
		)
	})
	should := require.New(t)
	err := Validate(TestObject{})
	should.NotNil(err)
	should.Contains(err.Error(), "myField")
	err = Validate(TestObject{1})
	should.Nil(err)
}
