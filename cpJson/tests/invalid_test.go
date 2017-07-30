package tests

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	"github.com/json-iterator/go"
	"errors"
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

func Test_stream_err(t *testing.T) {
	should := require.New(t)
	stream := jsoniter.NewStream(jsoniter.ConfigDefault, &faultyWriter{}, 512)
	srcBytes := []byte{}
	for i := 0; i < 2048; i++ {
		srcBytes = append(srcBytes, 'A')
	}
	src := string(srcBytes)
	should.NotNil(plz.Copy(stream, src))
}

type faultyWriter struct {
}

func (writer *faultyWriter) Write([]byte) (int, error) {
	return 0, errors.New("faulty")
}