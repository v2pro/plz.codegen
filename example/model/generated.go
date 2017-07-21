
package model
	
import "github.com/v2pro/wombat/generic"
func init() {
generic.RegisterExpandedFunc("MaxByItselfForPlz_T_int",MaxByItselfForPlz_T_int)
generic.RegisterExpandedFunc("MaxByFieldForPlz_F_Score_T_model__User",MaxByFieldForPlz_F_Score_T_model__User)
generic.RegisterExpandedFunc("New_Pair_I_model__IntStringPair",New_Pair_I_model__IntStringPair)}
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
func CompareByField_F_Score_T_model__User(val1 User,val2 User)( int){


return CompareByItself_T_int(val1.Score, val2.Score)
}
func MaxByFieldForPlz_F_Score_T_model__User(vals []interface{})( interface{}){

currentMax := vals[0].(User)
for i := 1; i < len(vals); i++ {
	typedVal := vals[i].(User)
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