
package model
import "unsafe"
import "fmt"
import "io"
import "github.com/v2pro/wombat/gen"

var ioEOF = io.EOF
var debugLog = fmt.Println

type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}
func init() {
	gen.RegisterFunc("Exported_Max_int", Exported_Max_int)
gen.RegisterFunc("Exported_Max_model__User_by_Score", Exported_Max_model__User_by_Score)
}
func Exported_obj_ptr_int(obj interface{}) unsafe.Pointer {
	return obj_ptr_int(obj)
}
func obj_ptr_int(obj interface{}) unsafe.Pointer {
	ptr := (*((*emptyInterface)(unsafe.Pointer(&obj)))).word
	
	return ptr
}

// generated from objPtr

func Exported_Compare_int(
	obj1 interface{},
	obj2 interface{}) int {
	// end of signature
	return Compare_int(
		(*(*int)(obj_ptr_int(obj1))),
		(*(*int)(obj_ptr_int(obj2))))
}

	func Compare_int(
		obj1 int,
		obj2 int) int {
		// end of signature
		if (obj1 < obj2) {
			return -1
		} else if (obj1 == obj2) {
			return 0
		} else {
			return 1
		}
	}


// generated from cmpSimpleValue


func Exported_Max_int(objs []interface{}) interface{} {
	currentMax := objs[0].(int)
	for i := 1; i < len(objs); i++ {
		typedObj := objs[i].(int)
		if Compare_int(typedObj, currentMax) > 0 {
			currentMax = typedObj
		}
	}
	return currentMax
}
func Max_int(objs []int) int {
	currentMax := objs[0]
	for i := 1; i < len(objs); i++ {
		if Compare_int(objs[i], currentMax) > 0 {
			currentMax = objs[i]
		}
	}
	return currentMax
}
// generated from maxSimpleValue

type model__User struct {
	Score int
}
func Exported_obj_ptr_model__User(obj interface{}) unsafe.Pointer {
	return obj_ptr_model__User(obj)
}
func obj_ptr_model__User(obj interface{}) unsafe.Pointer {
	ptr := (*((*emptyInterface)(unsafe.Pointer(&obj)))).word
	
	return ptr
}

// generated from objPtr

func Exported_Compare_model__User_by_Score(
	obj1 interface{},
	obj2 interface{}) int {
	// end of signature
	return Compare_model__User_by_Score(
		(*(*model__User)(obj_ptr_model__User(obj1))),
		(*(*model__User)(obj_ptr_model__User(obj2))))
}

	
	
	func Compare_model__User_by_Score(
		obj1 model__User,
		obj2 model__User) int {
		// end of signature
		return Compare_int(obj1.Score, obj2.Score)
	}

// generated from cmpStructByField


func Exported_Max_model__User_by_Score(objs []interface{}) interface{} {
	currentMaxObj := objs[0]
	for i := 1; i < len(objs); i++ {
		currentMax := (*(*model__User)(obj_ptr_model__User(currentMaxObj)))
		elem := (*(*model__User)(obj_ptr_model__User(objs[i])))
		if Compare_model__User_by_Score(elem, currentMax) > 0 {
			currentMaxObj = objs[i]
		}
	}
	return currentMaxObj
}
// generated from maxStructByField
