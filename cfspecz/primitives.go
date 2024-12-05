package cfspecz

// IsValidPrimitiveType validates a primitive type.
func IsValidPrimitiveType(primitiveType string) bool {
	switch primitiveType {
	case "Boolean", "Double", "Integer", "Json", "Long", "String", "Timestamp":
		return true
	default:
		return false
	}
}
