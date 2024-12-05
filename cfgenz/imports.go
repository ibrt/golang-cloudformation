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

func (ic *importsCollector) collectImports(goPackages ...string) {
	if ic != nil {
		for _, goPackage := range goPackages {
			ic.imports[goPackage] = struct{}{}
		}
	}
}

func (ic *importsCollector) toSlice() []string {
	return memz.GetSortedMapKeys(ic.imports, cmp.Less)
}
