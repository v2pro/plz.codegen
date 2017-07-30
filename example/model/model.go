package model

type IntStringPair interface {
	First() int
	SetFirst(val int)
	Second() string
	SetSecond(val string)
}

type User struct {
	Score int
}

type UserProperties struct {
	City string
	Age  int
}
type User2 struct {
	FirstName  string
	LastName   string
	Tags       []int
	Properties *UserProperties
}
type UserInfo struct {
	FirstName  *string
	LastName   *string
	Tags       []int
	Properties map[string]interface{}
}