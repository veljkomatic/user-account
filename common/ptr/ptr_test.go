package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CustomStruct struct {
	Field string
}

func TestFrom(t *testing.T) {
	// Test with int
	intVal := 5
	intPtr := From(intVal)
	assert.Equal(t, intVal, *intPtr, "From should return a pointer to the original int value")

	// Test with string
	strVal := "test"
	strPtr := From(strVal)
	assert.Equal(t, strVal, *strPtr, "From should return a pointer to the original string value")

	// Test with custom struct
	structVal := CustomStruct{Field: "value"}
	structPtr := From(structVal)
	assert.Equal(t, structVal, *structPtr, "From should return a pointer to the original struct value")
}

func TestValue(t *testing.T) {
	// Test with non-nil pointer
	intPtr := new(int)
	*intPtr = 10
	assert.Equal(t, 10, Value(intPtr), "Value should return the dereferenced value of a non-nil pointer")

	// Test with nil pointer and no default value
	var nilIntPtr *int
	assert.Equal(t, 0, Value(nilIntPtr), "Value should return the zero value for a nil pointer with no default value")

	// Test with nil pointer and default value
	defaultVal := 15
	assert.Equal(t, defaultVal, Value(nilIntPtr, defaultVal), "Value should return the provided default value for a nil pointer")

	// Test with string
	strPtr := new(string)
	*strPtr = "hello"
	assert.Equal(t, "hello", Value(strPtr), "Value should return the dereferenced value of a string pointer")
}
