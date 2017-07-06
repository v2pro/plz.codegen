package wombat
//
//import (
//	"reflect"
//	"fmt"
//	"unsafe"
//	"github.com/v2pro/plz/lang"
//	"github.com/v2pro/plz/lang/tagging"
//	"github.com/v2pro/plz/util"
//)
//
//func Validate(obj interface{}) error {
//	validator, err := validatorOfType(lang.AccessorOf(reflect.TypeOf(obj)), nil)
//	if err != nil {
//		return err
//	}
//	collector := newCollector()
//	ptr := extractPtr(obj)
//	collector.Enter("root", ptr)
//	err = validator.Validate(collector, ptr)
//	if err != nil {
//		return err
//	}
//	collector.Leave()
//	return collector.result()
//}
//
//func validatorOfType(accessor lang.Accessor, fieldTags map[string]interface{}) (util.Validator, error) {
//	switch accessor.Kind() {
//	case lang.Struct:
//		return validatorOfStruct(accessor)
//	case lang.Int:
//		return validatorOfInt(accessor, fieldTags)
//	default:
//		return nil, fmt.Errorf("do not know how to validate: %v", accessor)
//	}
//}
//
//func validatorOfStruct(accessor lang.Accessor) (util.Validator, error) {
//	fields := []structValidatorField{}
//	for i := 0; i < accessor.NumField(); i++ {
//		field := accessor.Field(i)
//		fieldTags := field.Tags()
//		valueValidator, err := validatorOfType(field.Accessor(), fieldTags)
//		if err != nil {
//			return nil, fmt.Errorf("field %s: %v", field.Name, err.Error())
//		}
//		fields = append(fields, structValidatorField{field, valueValidator})
//	}
//	return &structValidator{fields}, nil
//}
//
//func validatorOfInt(accessor lang.Accessor, fieldTags tagging.FieldTags) (util.Validator, error) {
//	if "required" == fieldTags["validate"] {
//		return &intRequiredValidator{accessor}, nil
//	}
//	return nil, nil
//}
//
//type structValidator struct {
//	fields []structValidatorField
//}
//
//type structValidatorField struct {
//	field          lang.StructField
//	valueValidator util.Validator
//}
//
//func (validator *structValidator) Validate(collector util.ResultCollector, obj interface{}) error {
//	for _, fieldValidator := range validator.fields {
//		fieldAccessor := fieldValidator.field.Accessor()
//		collector.Enter(fieldValidator.field.Name, unsafe.Pointer(fieldAccessor.AddressOf(obj)))
//		fieldValidator.valueValidator.Validate(collector, obj)
//		collector.Leave()
//	}
//	return nil
//}
//
//type intRequiredValidator struct {
//	accessor lang.Accessor
//}
//
//func (validator *intRequiredValidator) Validate(collector util.ResultCollector, obj interface{}) error {
//	val := validator.accessor.Int(obj)
//	if val == 0 {
//		collector.CollectError(fmt.Errorf("int is zero"))
//	}
//	return nil
//}
//
//func extractPtr(val interface{}) unsafe.Pointer {
//	return (*((*emptyInterface)(unsafe.Pointer(&val)))).word
//}
//
//// emptyInterface is the header for an interface{} value.
//type emptyInterface struct {
//	typ  unsafe.Pointer
//	word unsafe.Pointer
//}
