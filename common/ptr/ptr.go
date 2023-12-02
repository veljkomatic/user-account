package ptr

// From returns a pointer value for any value you passed in.
// Because of type safety you will have to cast *T to (*ConcreteType)
func From[T any](val T) *T {
	return &val
}

func Value[T any](ptr *T, optionalDefault ...T) T {
	if ptr == nil {
		if len(optionalDefault) > 0 {
			return optionalDefault[0]
		}
		return *(new(T))
	}
	return *ptr
}
