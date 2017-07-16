package invalid

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	"github.com/json-iterator/go"
	_ "github.com/v2pro/wombat/cpJson"
)

func Test_invalid_int(t *testing.T) {
	should := require.New(t)
	dst := int(0)
	src := `"100"`
	should.NotNil(jsonCopy(&dst, src))
}

func Test_indirect_invalid_int(t *testing.T) {
	should := require.New(t)
	var dst int
	var pDst interface{} = &dst
	src := `"100"`
	should.NotNil(jsonCopy(&pDst, src))
}

func jsonCopy(dst interface{}, srcJson string) error {
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, srcJson)
	return plz.Copy(dst, iter)
}
