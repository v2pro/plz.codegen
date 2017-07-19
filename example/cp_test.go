package example

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	_ "github.com/v2pro/wombat/cp"
	"testing"
)

type MyInterface interface{
	Hello()
}

func Test_interface(t *testing.T) {
	obj := new(MyInterface)
	(*obj).Hello()
}

func Test_copy(t *testing.T) {
	should := require.New(t)

	type UserProperties struct {
		City string
		Age  int
	}
	type User struct {
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
	userInfo := UserInfo{}
	plz.Copy(&userInfo, User{
		FirstName: "A",
		LastName:  "B",
		Tags:      []int{1, 2, 3},
		Properties: &UserProperties{
			"C",
			30,
		},
	})
	should.Equal("A", *userInfo.FirstName)
	should.Equal("B", *userInfo.LastName)
	should.Equal([]int{1, 2, 3}, userInfo.Tags)
	should.Equal("C", userInfo.Properties["City"])
	should.Equal(30, userInfo.Properties["Age"])
}
