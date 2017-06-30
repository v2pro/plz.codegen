package tf

import (
	"testing"
	"github.com/json-iterator/go"
	"github.com/v2pro/plz"
	"reflect"
	"github.com/stretchr/testify/require"
	_ "github.com/v2pro/wombat/jsonacc"
	_ "github.com/v2pro/plz/acc/native"
)

func Test_transform_int(t *testing.T) {
	should := require.New(t)
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, "1")
	iterAccessor := plz.AccessorOf(reflect.TypeOf(iter))
	var v *int
	nativeAccessor := plz.AccessorOf(reflect.TypeOf(v))
	transformedAccessor, err := Transform(iterAccessor, nativeAccessor)
	should.Nil(err)
	should.Equal(1, transformedAccessor.Int(iter))
}
