package acc

import (
	"fmt"
	"github.com/v2pro/plz/lang"
	"reflect"
	"unsafe"
)

func init() {
	lang.AccessorProviders = append(lang.AccessorProviders, accessorOfNativeType)
}

func objAcc(obj interface{}) lang.Accessor {
	return accessorOfNativeType(reflect.TypeOf(obj))
}

func objPtr(obj interface{}) unsafe.Pointer {
	return lang.AddressOf(obj)
}

func accessorOfNativeType(typ reflect.Type) lang.Accessor {
	switch typ.Kind() {
	case reflect.Ptr:
		elemType := typ.Elem()
		switch elemType.Kind() {
		case reflect.Interface:
			return &ptrVariantAccessor{
				variantAccessor{
					NoopAccessor: lang.NoopAccessor{"ptrVariantAccessor"},
					typ:          typ,
				},
			}
		case reflect.Map:
			fallthrough
		case reflect.Ptr:
			return &ptrPtrAccessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{"ptrPtrAccessor"},
				valueAccessor: accessorOfNativeType(elemType),
			}}
		case reflect.Int:
			return &ptrIntAccessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{"ptrIntAccessor"},
				valueAccessor: accessorOfNativeType(elemType),
			}}
		case reflect.Float64:
			return &ptrFloat64Accessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{"ptrFloat64Accessor"},
				valueAccessor: accessorOfNativeType(elemType),
			}}
		case reflect.String:
			return &ptrStringAccessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{"ptrStringAccessor"},
				valueAccessor: accessorOfNativeType(elemType),
			}}
		case reflect.Slice:
			return &ptrSliceAccessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{"ptrSliceAccessor"},
				valueAccessor: accessorOfNativeType(elemType),
			}}
		case reflect.Array:
			return &ptrArrayAccessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{"ptrArrayAccessor"},
				valueAccessor: accessorOfNativeType(elemType),
			}}
		case reflect.Struct:
			structAccessor := accessorOfNativeType(elemType).(*structAccessor)
			for i, field := range structAccessor.fields {
				field.accessor = accessorOfNativeType(reflect.PtrTo(elemType.Field(i).Type))
			}
			return &ptrStructAccessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{"ptrStructAccessor"},
				valueAccessor: structAccessor,
			}}
			//case reflect.Interface:
			//	return &ptrVariantAccessor{ptrAccessor{
			//		NoopAccessor:  lang.NoopAccessor{"ptrVariantAccessor"},
			//		valueAccessor: accessorOfNativeType(elemType),
			//	}}
			//case reflect.Interface:
			//	return &variantAccessor{
			//		lang.NoopAccessor{"variantAccessor"}}
			//case reflect.Struct:
			//	fallthrough
			//case reflect.Slice:
			//	fallthrough
			//case reflect.Array:
			//	return accessorOfNativeType(elemType)
			//
		}
	case reflect.Int:
		return &intAccessor{
			NoopAccessor: lang.NoopAccessor{"intAccessor"},
			typ:          typ,
		}
	case reflect.Float64:
		return &float64Accessor{
			NoopAccessor: lang.NoopAccessor{"float64Accessor"},
			typ:          typ,
		}
	case reflect.String:
		return &stringAccessor{
			NoopAccessor: lang.NoopAccessor{"stringAccessor"},
			typ:          typ,
		}
	case reflect.Struct:
		return accessorOfStruct(typ)
	case reflect.Map:
		templateEmptyInterface := castToEmptyInterface(reflect.New(typ).Elem().Interface())
		if typ.Elem().Kind() == reflect.Interface {
			return &mapInterfaceAccessor{
				mapAccessor{
					NoopAccessor:           lang.NoopAccessor{"mapAccessor"},
					typ:                    typ,
					templateEmptyInterface: templateEmptyInterface,
				},
			}
		}
		return &mapAccessor{
			NoopAccessor:           lang.NoopAccessor{"mapAccessor"},
			typ:                    typ,
			templateEmptyInterface: templateEmptyInterface,
		}
	case reflect.Slice:
		return &sliceAccessor{
			NoopAccessor: lang.NoopAccessor{"sliceAccessor"},
			elemAcc:      lang.AccessorOf(reflect.PtrTo(typ.Elem())),
			typ:          typ,
		}
	case reflect.Array:
		return &arrayAccessor{
			NoopAccessor: lang.NoopAccessor{"arrayAccessor"},
			elemAcc:      lang.AccessorOf(reflect.PtrTo(typ.Elem())),
			typ:          typ,
		}
	case reflect.Interface:
		return &variantAccessor{
			NoopAccessor: lang.NoopAccessor{"variantAccessor"},
			typ:          typ,
		}
	}
	panic(fmt.Sprintf("do not support: %v of kind %v", typ, typ.Kind()))
}
