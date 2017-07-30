package model

type TypeA struct {
	Int int
	String string
}

type TypeB struct {
	Field *int
}

type TypeC struct {
	Slice []int
	Array [3]int
	Map	map[string]int
	Struct TypeD
}

type TypeD struct {
	Int int
	String string
}

type TypeE struct {
	Field1 string
	Field3 string
}