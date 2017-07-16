package model

import (
	"github.com/v2pro/wombat"
	"github.com/v2pro/wombat/fp/maxSimpleValue"
	"github.com/v2pro/wombat/fp/maxStructByField"
	"reflect"
)

func init() {
	wombat.Declare(maxSimpleValue.F,
		"T", wombat.Int)
	wombat.Declare(maxStructByField.F,
		"T", reflect.TypeOf(User{}),
		"F", "Score")
}

type User struct {
	Score int
}