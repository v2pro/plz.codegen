package wombat

import (
	"fmt"
	"github.com/v2pro/plz/util"
)

func Example_copy_go_object() {
	type User struct {
		FirstName string
		LastName  string
	}
	type UserInfo struct {
		FirstName *string
		LastName  *string
	}
	userInfo := UserInfo{}
	util.Copy(&userInfo, User{"A", "B"})
	fmt.Println(*userInfo.FirstName)
	fmt.Println(*userInfo.LastName)
	// Output:
	// A
	// B
}

//func Example_encode_json() {
//	type User struct {
//		FirstName string `json:"first_name"`
//		LastName  string `json:"last_name"`
//		Tags      []int `json:"tags"`
//	}
//	tagging.Define(new(User), "codec", "json")
//	output := []byte{}
//	util.Copy(&output, User{"A", "B", []int{1, 2, 3}})
//	fmt.Println(string(output))
//	// Output:
//}
