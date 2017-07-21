package model

import (
	"reflect"
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/fp/compare"
)

func init() {
	generic.DeclareFunc(compare.ByItself,
		"T", generic.Int)
	generic.DeclareFunc(compare.ByField,
		"T", reflect.TypeOf(User{}),
		"F", "Score")
}

type User struct {
	Score int
}