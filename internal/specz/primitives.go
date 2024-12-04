package specz

func isValidPrimitiveType(primitiveType string) bool {
	switch primitiveType {
	case "Boolean", "Double", "Integer", "Json", "Long", "String", "Timestamp":
		return true
	default:
		return false
	}
}
