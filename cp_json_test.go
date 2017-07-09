package wombat

import (
	"github.com/v2pro/plz/lang/tagging"
	"fmt"
	_ "github.com/v2pro/wombat/cp_json"
	"github.com/v2pro/plz"
)

func Example_encode_json() {
	type User struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Tags      []int `json:"tags"`
	}
	tagging.Define(new(User), "codec", "json")

	output := []byte{}
	plz.Copy(&output, User{"A", "B", []int{1, 2, 3}})
	fmt.Println(string(output))
	// Output:
	// {"first_name":"A","last_name":"B","tags":[1,2,3]}
}
