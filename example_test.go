package wombat

import (
	"github.com/v2pro/plz/util"
	"fmt"
)

func Example_copy_go_object() {
	type User struct {
		FirstName string
		LastName string
	}
	type UserInfo struct {
		FirstName *string
		LastName *string
	}
	userInfo := UserInfo{}
	util.Copy(&userInfo, User{"A", "B"})
	fmt.Println(*userInfo.FirstName)
	fmt.Println(*userInfo.LastName)
	// Output:
	// A
	// B
}
