package wombat

import (
	"reflect"
	"fmt"
	"unsafe"
	"github.com/v2pro/plz/tags"
)

func Validate(obj interface{}) error {
	validator, err := validatorOfType(reflect.TypeOf(obj), nil)
	if err != nil {
		return err
	}
	collector := newCollector()
	ptr := extractPtr(obj)
	collector.Enter("root", ptr)
	err = validator.Validate(collector, ptr)
	if err != nil {
		return err
	}
	collector.Leave()
	return collector.result()
}

type Validator interface {
	Validate(collector ResultCollector, ptr unsafe.Pointer) error
}

func validatorOfType(typ reflect.Type, fieldTags tags.FieldTags) (Validator, error) {
	switch typ.Kind() {
	case reflect.Struct:
		return validatorOfStruct(typ)
	case reflect.Int:
		return validatorOfInt(typ, fieldTags)
	default:
		return nil, fmt.Errorf("do not know how to validate: %v", typ)
	}
}

func validatorOfStruct(typ reflect.Type) (Validator, error) {
	fields := []structValidatorField{}
	structTags := tags.Get(typ)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldTags := structTags.Fields[field.Name]
		valueValidator, err := validatorOfType(field.Type, fieldTags)
		if err != nil {
			return nil, fmt.Errorf("field %s: %v", field.Name, err.Error())
		}
		fields = append(fields, structValidatorField{field.Offset, field.Name, valueValidator})
	}
	return &structValidator{fields}, nil
}

func validatorOfInt(typ reflect.Type, fieldTags tags.FieldTags) (Validator, error) {
	if "required" == fieldTags["validate"] {
		return &intRequiredValidator{}, nil
	}
	return nil, nil
}

type structValidator struct {
	fields []structValidatorField
}

type structValidatorField struct {
	offset         uintptr
	fieldName      string
	valueValidator Validator
}

func (validator *structValidator) Validate(collector ResultCollector, ptr unsafe.Pointer) error {
	for _, field := range validator.fields {
		fieldPtr := unsafe.Pointer(uintptr(ptr) + field.offset)
		collector.Enter(field.fieldName, fieldPtr)
		field.valueValidator.Validate(collector, fieldPtr)
		collector.Leave()
	}
	return nil
}

type intRequiredValidator struct {
}

func (validator *intRequiredValidator) Validate(collector ResultCollector, ptr unsafe.Pointer) error {
	val := *((*int)(ptr))
	if val == 0 {
		collector.CollectError(fmt.Errorf("int is zero"))
	}
	return nil
}

func extractPtr(val interface{}) unsafe.Pointer {
	return (*((*emptyInterface)(unsafe.Pointer(&val)))).word
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}
