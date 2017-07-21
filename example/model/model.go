package model

import (
	"reflect"
	"github.com/v2pro/wombat/generic"
	"github.com/v2pro/wombat/fp/max"
	"github.com/v2pro/wombat/container/pair"
)

func init() {
	generic.DeclareFunc(max.ByItselfForPlz,
		"T", generic.Int)
	generic.DeclareFunc(max.ByFieldForPlz,
		"T", reflect.TypeOf(User{}),
		"F", "Score")
	generic.DeclareStruct(pair.Pair, reflect.TypeOf(new(IntStringPair)).Elem())
}

type IntStringPair interface {
	First() int
	SetFirst(val int)
	Second() string
	SetSecond(val string)
}

type User struct {
	Score int
}