package utils

func ToPtr[T any](item T) *T {
	return &item
}

func FromPtr[T any](itemPtr *T) T {
	var defaultVal T
	return FromPtrWithDefault(itemPtr, defaultVal)
}

func FromPtrWithDefault[T any](itemPtr *T, def T) T {
	if itemPtr != nil {
		return *itemPtr
	}

	return def
}
