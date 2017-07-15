package test

type typeA struct {
	Int int
	String string
}

type typeForTest struct {
	Slice []int
	Array [3]int
	Map	map[string]int
	Struct typeA
}