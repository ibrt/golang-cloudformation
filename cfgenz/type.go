package cfgenz

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/memz"
	"github.com/ibrt/golang-utils/tplz"

	"github.com/ibrt/golang-cloudformation/cfschemaz"
	"github.com/ibrt/golang-cloudformation/cfspecz"
)

// Type describes either a top-level resource type or a structured type,
type Type struct {
	Attributes map[string]*Attribute
	Properties map[string]*Property

	g       *Generator
	specT   *cfspecz.Type
	schemaT *cfschemaz.Type
}

func newType(g *Generator, specT *cfspecz.Type, schemaT *cfschemaz.Type) *Type {
	if schemaT == nil {
		if specT.IsTopLevelResourceType {
			fmt.Println("TopLevelResourceType:", specT.Name)
		} else {
			fmt.Println("StructuredType:", specT.Name)
		}
	}

	t := &Type{
		g:       g,
		specT:   specT,
		schemaT: schemaT,
	}

	t.Attributes = memz.TransformMapValues(specT.Attributes, func(_ string, specA *cfspecz.Attribute) *Attribute {
		return newAttribute(t, specA)
	})

	t.Properties = memz.TransformMapValues(specT.Properties, func(_ string, specP *cfspecz.Property) *Property {
		return newProperty(t, specP)
	})

	return t
}

// GetRelatedTopLevelResourceTypeName returns the top-level resource type name related to this type, that is, Name if
// IsTopLevelResourceType = true, or the portion of Name before the first "." otherwise.
func (t *Type) GetRelatedTopLevelResourceTypeName() string {
	return t.specT.GetRelatedTopLevelResourceTypeName()
}

// GetRelatedStructuredTypeName converts the value of a Type or ItemType field of a property or attribute in this type
// to its fully qualified structured type name.
func (t *Type) GetRelatedStructuredTypeName(unqualifiedStructuredTypeName string) string {
	return t.specT.GetRelatedStructuredTypeName(unqualifiedStructuredTypeName)
}

// GoName returns the Go name for this type.
func (t *Type) GoName() string {
	structName := strings.ReplaceAll(t.specT.Name, "::", "_")
	return strings.ReplaceAll(structName, ".", "_")
}

// GoComment returns a Go comment for this type.
func (t *Type) GoComment() string {
	return NewComment().
		AddLinef("%v is a binding for %v.", t.GoName(), t.Name()).
		MaybeAddLink("Documentation", t.specT.Documentation).
		String()
}

// GoPackageName returns the Go package name for this type.
func (t *Type) GoPackageName() string {
	return strings.ToLower(strings.Join(
		strings.Split(t.GetRelatedTopLevelResourceTypeName(), "::")[0:2],
		"_"))
}

// GoFileName returns the Go file name for this type.
func (t *Type) GoFileName() string {
	return strings.ToLower(t.GoName()) + ".go"
}

// GoImports returns the Go imports for this type.
func (t *Type) GoImports() []string {
	ic := newImportsCollector()
	ic.collectImports(t.g.o.TypeTemplateGetImplicitGoImports(t)...)

	for _, ga := range t.Attributes {
		ga.collectImports(ic)
	}

	for _, gp := range t.Properties {
		gp.collectImports(ic)
	}

	return ic.toSlice()
}

// GoSupportBasePackage returns the base package for the support library.
func (t *Type) GoSupportBasePackage() string {
	return t.g.o.getGoSupportBasePackage()
}

// IsTopLevelResourceType returns true if this is a top-level resource type, false if it is a structured type.
func (t *Type) IsTopLevelResourceType() bool {
	return t.specT.IsTopLevelResourceType
}

// Name returns the CloudFormation name for this type.
func (t *Type) Name() string {
	return t.specT.Name
}

func (t *Type) generate(outDirPath string) error {
	buf, err := tplz.ExecuteGo(t.g.o.TypeTemplate, t)
	if err != nil {
		return errorz.Wrap(err)
	}

	outDirPath = filepath.Join(outDirPath, t.GoPackageName())
	if err := os.MkdirAll(outDirPath, 0777); err != nil {
		return errorz.Wrap(err)
	}

	return errorz.MaybeWrap(os.WriteFile(filepath.Join(outDirPath, t.GoFileName()), buf, 0666))
}
