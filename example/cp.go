package example

import (
	"github.com/stretchr/testify/require"
	"github.com/v2pro/plz"
	"testing"
	_ "github.com/v2pro/wombat/cp"
	"github.com/v2pro/wombat/example/model"
	"github.com/v2pro/wombat/generic"
)

func init() {
	generic.Declare(func() {
		plz.Copy(&model.UserInfo{}, model.User2{})
		strVal := ""
		plz.Copy(&strVal, strVal)
		intVal := int(0)
		plz.Copy(&intVal, intVal)
	})
}

func Demo_copy(t *testing.T) {
	should := require.New(t)

	userInfo := model.UserInfo{}
	plz.Copy(&userInfo, model.User2{
		FirstName: "A",
		LastName:  "B",
		Tags:      []int{1, 2, 3},
		Properties: &model.UserProperties{
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
