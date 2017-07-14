package diff

import (
	"reflect"
	"strings"
	"fmt"
)

func Diff(obj1 interface{}, obj2 interface{}) []string {
	return doDiff([]string{}, obj1, obj2)
}

func doDiff(path []string, obj1 interface{}, obj2 interface{}) (logs []string) {
	switch reflect.TypeOf(obj1).Kind() {
	case reflect.Int:
		if reflect.ValueOf(obj1).Int() != reflect.ValueOf(obj2).Int() {
			logs = append(logs, fmt.Sprintf("%s: %v <=> %v", strings.Join(path, "."), obj1, obj2))
		}
	default:
		panic(strings.Join(path, ".") +": do not know how to diff")
	}
	return
}
