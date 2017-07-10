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
	return accessorOfNativeType(reflect.TypeOf(obj), "")
}

func objPtr(obj interface{}) unsafe.Pointer {
	return lang.AddressOf(obj)
}

func accessorOfNativeType(typ reflect.Type, tagName string) lang.Accessor {
	switch typ.Kind() {
	case reflect.Ptr:
		elemType := typ.Elem()
		switch elemType.Kind() {
		case reflect.Interface:
			return &ptrVariantAccessor{
				variantAccessor{
					NoopAccessor: lang.NoopAccessor{tagName,"ptrVariantAccessor"},
					typ:          typ,
				},
			}
		case reflect.Map:
			fallthrough
		case reflect.Ptr:
			return &ptrPtrAccessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{tagName,"ptrPtrAccessor"},
				valueAccessor: accessorOfNativeType(elemType, tagName),
			}}
		case reflect.Int:
			return &ptrIntAccessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{tagName,"ptrIntAccessor"},
				valueAccessor: accessorOfNativeType(elemType, tagName),
			}}
		case reflect.Int8:
			return &ptrInt8Accessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{tagName,"ptrInt8Accessor"},
				valueAccessor: accessorOfNativeType(elemType, tagName),
			}}
		case reflect.Float64:
			return &ptrFloat64Accessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{tagName,"ptrFloat64Accessor"},
				valueAccessor: accessorOfNativeType(elemType, tagName),
			}}
		case reflect.String:
			return &ptrStringAccessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{tagName,"ptrStringAccessor"},
				valueAccessor: accessorOfNativeType(elemType, tagName),
			}}
		case reflect.Slice:
			return &ptrSliceAccessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{tagName,"ptrSliceAccessor"},
				valueAccessor: accessorOfNativeType(elemType, tagName),
			}}
		case reflect.Array:
			return &ptrArrayAccessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{tagName,"ptrArrayAccessor"},
				valueAccessor: accessorOfNativeType(elemType, tagName),
			}}
		case reflect.Struct:
			structAccessor := accessorOfNativeType(elemType, tagName).(*structAccessor)
			return &ptrStructAccessor{ptrAccessor{
				NoopAccessor:  lang.NoopAccessor{tagName,"ptrStructAccessor"},
				valueAccessor: structAccessor,
			}}
		}
	case reflect.Int:
		return &intAccessor{
			NoopAccessor: lang.NoopAccessor{tagName,"intAccessor"},
			typ:          typ,
		}
	case reflect.Int8:
		return &int8Accessor{
			NoopAccessor: lang.NoopAccessor{tagName,"int8Accessor"},
			typ:          typ,
		}
	case reflect.Float64:
		return &float64Accessor{
			NoopAccessor: lang.NoopAccessor{tagName,"float64Accessor"},
			typ:          typ,
		}
	case reflect.String:
		return &stringAccessor{
			NoopAccessor: lang.NoopAccessor{tagName,"stringAccessor"},
			typ:          typ,
		}
	case reflect.Struct:
		return accessorOfStruct(typ, tagName)
	case reflect.Map:
		return accessorOfMap(typ, tagName)
	case reflect.Slice:
		return &sliceAccessor{
			NoopAccessor: lang.NoopAccessor{tagName,"sliceAccessor"},
			elemAcc:      lang.AccessorOf(reflect.PtrTo(typ.Elem()), tagName),
			typ:          typ,
		}
	case reflect.Array:
		return &arrayAccessor{
			NoopAccessor: lang.NoopAccessor{tagName,"arrayAccessor"},
			elemAcc:      lang.AccessorOf(reflect.PtrTo(typ.Elem()), tagName),
			typ:          typ,
		}
	case reflect.Interface:
		return &variantAccessor{
			NoopAccessor: lang.NoopAccessor{tagName,"variantAccessor"},
			typ:          typ,
		}
	}
	panic(fmt.Sprintf("do not support: %v of kind %v", typ, typ.Kind()))
}
