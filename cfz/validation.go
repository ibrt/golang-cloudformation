package cfz

// Validation describes a value.
type Validation struct {
	Single *ValidationSingle
	Array  *ValidationArray
}

// ValidationString describes a string value.
type ValidationString struct {
	MinLength *int
	MaxLength *int
	Pattern   *string
	Format    *string
	Enum      []string
}

// ValidationNumber describes a number value.
type ValidationNumber[T float64 | int64] struct {
	Minimum    *T
	Maximum    *T
	MultipleOf *T
	Enum       []T
}

// ValidationStructuredType describe a structured type value.
type ValidationStructuredType struct {
	StructuredTypeName string
}

// ValidationSingle describes a single value.
type ValidationSingle struct {
	String     *ValidationString
	Int64      *ValidationNumber[int64]
	Float64    *ValidationNumber[float64]
	Structured *ValidationStructuredType
}

// ValidationArray describes an array.
type ValidationArray struct {
	MinItems       *int
	MaxItems       *int
	InsertionOrder *bool
	UniqueItems    *bool
	Items          *Validation
}
