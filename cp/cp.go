package cp

import (
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
	_ "github.com/v2pro/wombat/acc"
)

func init() {
	util.CopierProviders = append(util.CopierProviders, provideCopier)
}

func provideCopier(dstAccessor, srcAccessor lang.Accessor) (util.Copier, error) {
	if srcAccessor.Kind() == lang.Variant {
		return &fromVariantCopier{
			srcAcc: srcAccessor,
			dstAcc: dstAccessor,
		}, nil
	}
	if dstAccessor.Kind() == lang.Variant {
		return &toVariantCopier{
			srcAcc: srcAccessor,
			dstAcc: dstAccessor,
		}, nil
	}
	if dstAccessor.Kind() == lang.Struct && dstAccessor.RandomAccessible() {
		switch srcAccessor.Kind() {
		case lang.Struct:
			return newStructToStructCopier(dstAccessor, srcAccessor)
		case lang.Map:
			return newMapToStructCopier(dstAccessor, srcAccessor)
		}
	}
	if dstAccessor.Kind() == lang.Map {
		switch srcAccessor.Kind() {
		case lang.Map:
			return newMapToMapCopier(dstAccessor, srcAccessor)
		case lang.Struct:
			return newStructToMapCopier(dstAccessor, srcAccessor)
		}
	}
	if srcAccessor.Kind() == dstAccessor.Kind() {
		switch srcAccessor.Kind() {
		case lang.Int:
			return &intCopier{
				srcAcc: srcAccessor,
				dstAcc: dstAccessor,
			}, nil
		case lang.Float64:
			return &float64Copier{
				srcAcc: srcAccessor,
				dstAcc: dstAccessor,
			}, nil
		case lang.String:
			return &stringCopier{
				srcAcc: srcAccessor,
				dstAcc: dstAccessor,
			}, nil
		case lang.Array:
			elemCopier, err := util.CopierOf(dstAccessor.Elem(), srcAccessor.Elem())
			if err != nil {
				return nil, err
			}
			return &arrayCopier{
				srcAcc:     srcAccessor,
				dstAcc:     dstAccessor,
				elemCopier: elemCopier,
			}, nil
		}
	}
	return nil, nil
}
