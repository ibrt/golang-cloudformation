package cfschemaz

// Property describes a property of either a top-level resource or structured type.
type Property struct {
	Name       string
	Validation *Validation
}

/*
func newProperty(t *Type, pUST *UnprocessedDefinition, name string) *Property {
	p := &Property{
		Name: name,
	}

	if pUST.Type == "string" {
		p.TypeString = &PropertyTypeString{
			MinLength: pUST.MinLength,
			MaxLength: pUST.MaxLength,
			Pattern:   pUST.Pattern,
			Enum:      pUST.Enum.AsString(),
		}

		return p
	}

	if pUST.Type == "integer" {
		p.TypeInteger = &PropertyTypeInteger{
			Enum: pUST.Enum.MustAsInt64(),
		}

		if pUST.Minimum != nil {
			v, err := pUST.Minimum.Int64()
			errorz.MaybeMustWrap(err)
			p.TypeInteger.Minimum = &v
		}

		if pUST.Maximum != nil {
			v, err := pUST.Maximum.Int64()
			errorz.MaybeMustWrap(err)
			p.TypeInteger.Maximum = &v
		}

		if pUST.MultipleOf != nil {
			v, err := pUST.MultipleOf.Int64()
			errorz.MaybeMustWrap(err)
			p.TypeInteger.MultipleOf = &v
		}

		return p
	}

	if pUST.Type == "number" {
		p.TypeNumber = &PropertyTypeNumber{
			Enum: pUST.Enum.MustAsFloat64(),
		}

		if pUST.Minimum != nil {
			v, err := pUST.Minimum.Float64()
			errorz.MaybeMustWrap(err)
			p.TypeNumber.Minimum = &v
		}

		if pUST.Maximum != nil {
			v, err := pUST.Maximum.Float64()
			errorz.MaybeMustWrap(err)
			p.TypeNumber.Maximum = &v
		}

		if pUST.MultipleOf != nil {
			v, err := pUST.MultipleOf.Float64()
			errorz.MaybeMustWrap(err)
			p.TypeNumber.MultipleOf = &v
		}

		return p
	}

	if pUST.Ref != nil {

	}

	return p
}
*/
