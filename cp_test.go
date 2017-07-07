package wombat

import (
	"fmt"
	_ "github.com/v2pro/wombat/cp"
	"github.com/v2pro/plz"
)

func Example_copy_struct() {
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
	fmt.Println(*userInfo.FirstName)
	fmt.Println(*userInfo.LastName)
	fmt.Println(userInfo.Tags)
	fmt.Println(userInfo.Properties["City"])
	fmt.Println(userInfo.Properties["Age"])
	// Output:
	// A
	// B
	// [1 2 3]
	// C
	// 30
}
