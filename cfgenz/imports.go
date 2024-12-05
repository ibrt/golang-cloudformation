package cfgenz

import (
	"cmp"

	"github.com/ibrt/golang-utils/memz"
)

type importsCollector struct {
	imports map[string]struct{}
}

func newImportsCollector() *importsCollector {
	return &importsCollector{
		imports: map[string]struct{}{},
	}
}

func (ic *importsCollector) collectImports(goPackages ...string) *importsCollector {
	if ic != nil {
		for _, goPackage := range goPackages {
			ic.imports[goPackage] = struct{}{}
		}
	}
	return ic
}

func (ic *importsCollector) maybeCollectImports(cond bool, goPackages ...string) *importsCollector {
	if ic != nil && cond {
		ic.collectImports(goPackages...)
	}
	return ic
}

func (ic *importsCollector) toSlice() []string {
	return memz.GetSortedMapKeys(ic.imports, cmp.Less)
}
