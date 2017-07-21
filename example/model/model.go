package model

import (
	"reflect"
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/fp/max"
)

func init() {
	generic.DeclareFunc(max.ByItselfForPlz,
		"T", generic.Int)
	generic.DeclareFunc(max.ByFieldForPlz,
		"T", reflect.TypeOf(User{}),
		"F", "Score")
}

type User struct {
	Score int
}