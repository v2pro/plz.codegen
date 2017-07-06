package wombat
//
//import (
//	"testing"
//	"github.com/stretchr/testify/require"
//	"github.com/v2pro/plz/lang/tagging"
//)
//
//func Test_int_required(t *testing.T) {
//	type TestObject struct {
//		myField int
//	}
//	tagging.Define(func(o *TestObject) tagging.Tags {
//		return tagging.D(tagging.S(),
//			tagging.F(&o.myField, "validate", "required"),
//		)
//	})
//	should := require.New(t)
//	err := Validate(TestObject{})
//	should.NotNil(err)
//	should.Contains(err.Error(), "myField")
//	err = Validate(TestObject{1})
//	should.Nil(err)
//}
