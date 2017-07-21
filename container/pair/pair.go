package pair

import "github.com/v2pro/wombat/generic"

var Pair = generic.DefineStruct("Pair").
	Source(`
{{ $T1 := .I | method "First" | returnType }}
{{ $T2 := .I | method "Second" | returnType }}

type {{.structName}} struct {
    first {{$T1|name}}
    second {{$T2|name}}
}

func (pair *{{.structName}}) SetFirst(val {{$T1|name}}) {
    pair.first = val
}

func (pair *{{.structName}}) First() {{$T1|name}} {
    return pair.first
}

func (pair *{{.structName}}) SetSecond(val {{$T2|name}}) {
    pair.second = val
}

func (pair *{{.structName}}) Second() {{$T2|name}} {
    return pair.second
}`)
