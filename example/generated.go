
package example
import "github.com/v2pro/wombat/generic"
import "io"
import "github.com/v2pro/wombat/example/model"
func init() {
generic.RegisterExpandedFunc("CopyAnythingForPlz_CT_gXD56NMK6FVXA26GM25XVH4ZILKLIKTKL_DT_ptr_model__UserInfo_ST_model__User2",CopyAnythingForPlz_CT_gXD56NMK6FVXA26GM25XVH4ZILKLIKTKL_DT_ptr_model__UserInfo_ST_model__User2)
generic.RegisterExpandedFunc("CopyAnythingForPlz_CT_gXD56NMK6FVXA26GM25XVH4ZILKLIKTKL_DT_ptr_string_ST_string",CopyAnythingForPlz_CT_gXD56NMK6FVXA26GM25XVH4ZILKLIKTKL_DT_ptr_string_ST_string)
generic.RegisterExpandedFunc("CopyAnythingForPlz_CT_gXD56NMK6FVXA26GM25XVH4ZILKLIKTKL_DT_ptr_int_ST_int",CopyAnythingForPlz_CT_gXD56NMK6FVXA26GM25XVH4ZILKLIKTKL_DT_ptr_int_ST_int)
generic.RegisterExpandedFunc("MaxByItselfForPlz_T_int",MaxByItselfForPlz_T_int)
generic.RegisterExpandedFunc("MaxByItselfForPlz_T_float64",MaxByItselfForPlz_T_float64)
generic.RegisterExpandedFunc("MaxByFieldForPlz_F_Score_T_model__User",MaxByFieldForPlz_F_Score_T_model__User)
generic.RegisterExpandedFunc("New_Pair_I_model__IntStringPair",New_Pair_I_model__IntStringPair)}
var copyDynamically func(interface{}, interface{}) error
var ioEOF = io.EOF
func CopySimpleValue_DT_ptr_string_ST_string(err *error,dst *string,src string){
if dst != nil {
	*dst = (string)(src)
}

}
func CopyAnything_DT_ptr_string_ST_string(err *error,dst *string,src string){


CopySimpleValue_DT_ptr_string_ST_string(err, dst, src)



}
func CopyIntoPtr_DT_ptr_ptr_string_ST_string(err *error,dst **string,src string){


defDst := *dst
if defDst == nil {
	
		defDst = new(string)
	
	CopyAnything_DT_ptr_string_ST_string(err, defDst, src)
	*dst = defDst
	return
}
CopyAnything_DT_ptr_string_ST_string(err, *dst, src)

}
func CopyAnything_DT_ptr_ptr_string_ST_string(err *error,dst **string,src string){


CopyIntoPtr_DT_ptr_ptr_string_ST_string(err, dst, src)



}
func CopySimpleValue_DT_ptr_int_ST_int(err *error,dst *int,src int){
if dst != nil {
	*dst = (int)(src)
}

}
func CopyAnything_DT_ptr_int_ST_int(err *error,dst *int,src int){


CopySimpleValue_DT_ptr_int_ST_int(err, dst, src)



}
func CopySliceToSlice_DT_ptr_slice_int_ST_slice_int(err *error,dst *[]int,src []int){

if src == nil {
	*dst = nil
	return
}
dstLen := len(*dst)
if len(src) < dstLen {
	dstLen = len(src)
}
for i := 0; i < dstLen; i++ {
	CopyAnything_DT_ptr_int_ST_int(err, &(*dst)[i], src[i])
}
defDst := *dst
for i := dstLen; i < len(src); i++ {
	newElem := new(int)
	CopyAnything_DT_ptr_int_ST_int(err, newElem, src[i])
	defDst = append(defDst, *newElem)
}
*dst = defDst

}
func CopyAnything_DT_ptr_slice_int_ST_slice_int(err *error,dst *[]int,src []int){


CopySliceToSlice_DT_ptr_slice_int_ST_slice_int(err, dst, src)



}
func CopyIntoInterface_DT_ptr_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_string(err *error,dst *interface {},src string){
if *dst == nil {
	newDst := new(string)
	newErr := copyDynamically(newDst, src)
	if newErr != nil && *err == nil {
		*err = newErr
	}
	*dst = *newDst
} else {
	newErr := copyDynamically(*dst, src)
	if newErr != nil && *err == nil {
		*err = newErr
	}
}
}
func CopyAnything_DT_ptr_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_string(err *error,dst *interface {},src string){


CopyIntoInterface_DT_ptr_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_string(err, dst, src)



}
func CopyIntoInterface_DT_ptr_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_int(err *error,dst *interface {},src int){
if *dst == nil {
	newDst := new(int)
	newErr := copyDynamically(newDst, src)
	if newErr != nil && *err == nil {
		*err = newErr
	}
	*dst = *newDst
} else {
	newErr := copyDynamically(*dst, src)
	if newErr != nil && *err == nil {
		*err = newErr
	}
}
}
func CopyAnything_DT_ptr_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_int(err *error,dst *interface {},src int){


CopyIntoInterface_DT_ptr_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_int(err, dst, src)



}
func CopyStructToMap_DT_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_model__UserProperties(err *error,dst map[string]interface {},src model.UserProperties){


	
	

	
	

var existingElem interface {}
var found bool

	existingElem, found = dst["City"]
	if found {
		CopyAnything_DT_ptr_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_string(err, &existingElem, src.City)
		dst["City"] = existingElem
	} else {
		newElem := new(interface {})
		CopyAnything_DT_ptr_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_string(err, newElem, src.City)
		dst["City"] = *newElem
	}

	existingElem, found = dst["Age"]
	if found {
		CopyAnything_DT_ptr_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_int(err, &existingElem, src.Age)
		dst["Age"] = existingElem
	} else {
		newElem := new(interface {})
		CopyAnything_DT_ptr_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_int(err, newElem, src.Age)
		dst["Age"] = *newElem
	}

}
func CopyAnything_DT_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_model__UserProperties(err *error,dst map[string]interface {},src model.UserProperties){


CopyStructToMap_DT_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_model__UserProperties(err, dst, src)



}
func CopyFromPtr_DT_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_ptr_model__UserProperties(err *error,dst map[string]interface {},src *model.UserProperties){

if src == nil {
	return
}
CopyAnything_DT_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_model__UserProperties(err, dst, *src)
}
func CopyAnything_DT_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_ptr_model__UserProperties(err *error,dst map[string]interface {},src *model.UserProperties){


CopyFromPtr_DT_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_ptr_model__UserProperties(err, dst, src)



}
func CopyIntoPtr_DT_ptr_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_ptr_model__UserProperties(err *error,dst *map[string]interface {},src *model.UserProperties){


	if src == nil {
		*dst = nil
		return
	}

defDst := *dst
if defDst == nil {
	
		defDst = map[string]interface {}{}
	
	CopyAnything_DT_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_ptr_model__UserProperties(err, defDst, src)
	*dst = defDst
	return
}
CopyAnything_DT_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_ptr_model__UserProperties(err, *dst, src)

}
func CopyAnything_DT_ptr_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_ptr_model__UserProperties(err *error,dst *map[string]interface {},src *model.UserProperties){


CopyIntoPtr_DT_ptr_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_ptr_model__UserProperties(err, dst, src)



}
func CopyStructToStruct_DT_ptr_model__UserInfo_ST_model__User2(err *error,dst *model.UserInfo,src model.User2){


	
	

	
	

	
	

	
	


	CopyAnything_DT_ptr_ptr_string_ST_string(err, &dst.FirstName, src.FirstName)

	CopyAnything_DT_ptr_ptr_string_ST_string(err, &dst.LastName, src.LastName)

	CopyAnything_DT_ptr_slice_int_ST_slice_int(err, &dst.Tags, src.Tags)

	CopyAnything_DT_ptr_map_string_to_gQPAJHSWGPSUIJTIJRIBX4GE44DBOXJ5R_ST_ptr_model__UserProperties(err, &dst.Properties, src.Properties)

}
func CopyAnything_DT_ptr_model__UserInfo_ST_model__User2(err *error,dst *model.UserInfo,src model.User2){


CopyStructToStruct_DT_ptr_model__UserInfo_ST_model__User2(err, dst, src)



}
func CopyAnythingForPlz_CT_gXD56NMK6FVXA26GM25XVH4ZILKLIKTKL_DT_ptr_model__UserInfo_ST_model__User2(theCopyDynamically func(interface {}, interface {}) error,dst interface{},src interface{})( error){

copyDynamically = theCopyDynamically
var err error
CopyAnything_DT_ptr_model__UserInfo_ST_model__User2(&err, dst.(*model.UserInfo), src.(model.User2))
return err
}
func CopyAnythingForPlz_CT_gXD56NMK6FVXA26GM25XVH4ZILKLIKTKL_DT_ptr_string_ST_string(theCopyDynamically func(interface {}, interface {}) error,dst interface{},src interface{})( error){

copyDynamically = theCopyDynamically
var err error
CopyAnything_DT_ptr_string_ST_string(&err, dst.(*string), src.(string))
return err
}
func CopyAnythingForPlz_CT_gXD56NMK6FVXA26GM25XVH4ZILKLIKTKL_DT_ptr_int_ST_int(theCopyDynamically func(interface {}, interface {}) error,dst interface{},src interface{})( error){

copyDynamically = theCopyDynamically
var err error
CopyAnything_DT_ptr_int_ST_int(&err, dst.(*int), src.(int))
return err
}
func CompareSimpleValue_T_int(val1 int,val2 int)( int){
if val1 < val2 {
	return -1
} else if val1 == val2 {
	return 0
} else {
	return 1
}
}
func CompareByItself_T_int(val1 int,val2 int)( int){

return CompareSimpleValue_T_int(val1, val2)
}
func MaxByItselfForPlz_T_int(vals []interface{})( interface{}){

currentMax := vals[0].(int)
for i := 1; i < len(vals); i++ {
	typedVal := vals[i].(int)
	if CompareByItself_T_int(typedVal, currentMax) > 0 {
		currentMax = typedVal
	}
}
return currentMax
}
func CompareSimpleValue_T_float64(val1 float64,val2 float64)( int){
if val1 < val2 {
	return -1
} else if val1 == val2 {
	return 0
} else {
	return 1
}
}
func CompareByItself_T_float64(val1 float64,val2 float64)( int){

return CompareSimpleValue_T_float64(val1, val2)
}
func MaxByItselfForPlz_T_float64(vals []interface{})( interface{}){

currentMax := vals[0].(float64)
for i := 1; i < len(vals); i++ {
	typedVal := vals[i].(float64)
	if CompareByItself_T_float64(typedVal, currentMax) > 0 {
		currentMax = typedVal
	}
}
return currentMax
}
func CompareByField_F_Score_T_model__User(val1 model.User,val2 model.User)( int){


return CompareByItself_T_int(val1.Score, val2.Score)
}
func MaxByFieldForPlz_F_Score_T_model__User(vals []interface{})( interface{}){

currentMax := vals[0].(model.User)
for i := 1; i < len(vals); i++ {
	typedVal := vals[i].(model.User)
	if CompareByField_F_Score_T_model__User(typedVal, currentMax) > 0 {
		currentMax = typedVal
	}
}
return currentMax
}



type Pair_I_model__IntStringPair struct {
    first int
    second string
}

func (pair *Pair_I_model__IntStringPair) SetFirst(val int) {
    pair.first = val
}

func (pair *Pair_I_model__IntStringPair) First() int {
    return pair.first
}

func (pair *Pair_I_model__IntStringPair) SetSecond(val string) {
    pair.second = val
}

func (pair *Pair_I_model__IntStringPair) Second() string {
    return pair.second
}
func New_Pair_I_model__IntStringPair()( interface{}){

return &Pair_I_model__IntStringPair{}
}