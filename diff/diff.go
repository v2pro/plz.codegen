package diff

import (
	"fmt"
	"reflect"
	"strings"
)

func Diff(obj1 interface{}, obj2 interface{}) []string {
	return doDiff([]string{}, obj1, obj2)
}

func doDiff(path []string, obj1 interface{}, obj2 interface{}) (logs []string) {
	switch reflect.TypeOf(obj1).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if reflect.ValueOf(obj1).Int() != reflect.ValueOf(obj2).Int() {
			logs = append(logs, fmt.Sprintf("%s: %v <=> %v", strings.Join(path, "."), obj1, obj2))
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if reflect.ValueOf(obj1).Uint() != reflect.ValueOf(obj2).Uint() {
			logs = append(logs, fmt.Sprintf("%s: %v <=> %v", strings.Join(path, "."), obj1, obj2))
		}
	case reflect.Float32, reflect.Float64:
		if reflect.ValueOf(obj1).Float() != reflect.ValueOf(obj2).Float() {
			logs = append(logs, fmt.Sprintf("%s: %v <=> %v", strings.Join(path, "."), obj1, obj2))
		}
	case reflect.String:
		if reflect.ValueOf(obj1).String() != reflect.ValueOf(obj2).String() {
			logs = append(logs, fmt.Sprintf("%s: %v <=> %v", strings.Join(path, "."), obj1, obj2))
		}
	default:
		panic(strings.Join(path, ".") + ": do not know how to diff " + reflect.TypeOf(obj1).String())
	}
	return
}
