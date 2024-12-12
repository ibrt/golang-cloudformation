package cfschemaz

// Type describes either a top-level resource type or a structured type.
type Type struct {
	IsTopLevelResourceType bool
	Name                   string
	Description            string
	Properties             map[string]*Property
}

/*
func (t *Type) process() {
	if t.IsTopLevelResourceType {
		for name, pUST := range t.utlr.Properties {
			t.Properties[name] = newProperty(t, pUST, name)
		}
	} else {
		for name, pUST := range t.ust.Properties {
			t.Properties[name] = newProperty(t, pUST, name)
		}
	}
}
*/
