
package test
import "github.com/v2pro/wombat/generic"
func init() {
generic.RegisterExpandedFunc("CopyAnythingForPlz_DT_ptr_test__DstType_ST_test__SrcType",CopyAnythingForPlz_DT_ptr_test__DstType_ST_test__SrcType)}
func CopySimpleValue_DT_ptr_test__DstType_ST_test__SrcType(err *error,dst *DstType,src SrcType)(){*dst = (DstType)(src)
}
func CopyAnything_DT_ptr_test__DstType_ST_test__SrcType(err *error,dst *DstType,src SrcType)(){


CopySimpleValue_DT_ptr_test__DstType_ST_test__SrcType(err, dst, src)
}
func CopyAnythingForPlz_DT_ptr_test__DstType_ST_test__SrcType(dst interface{},src interface{})( error){

var err error
CopyAnything_DT_ptr_test__DstType_ST_test__SrcType(&err, dst.(*DstType), src.(SrcType))
return err
}