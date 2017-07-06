package wombat

import (
	_ "github.com/v2pro/plz_native_accessor"
	"github.com/v2pro/plz/lang"
	"github.com/v2pro/plz/util"
)

func init() {
	util.CopierProviders = append(util.CopierProviders, provideCopier)
}

func provideCopier(dstAccessor, srcAccessor lang.Accessor) (util.Copier, error) {
	if srcAccessor.Kind() == dstAccessor.Kind() {
		switch srcAccessor.Kind() {
		case lang.Int:
			return &intCopier{
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
