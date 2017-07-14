package cpVal

type fromType struct {
	Field1 int
	Field2 struct {
		Field1 int
		Field2 string
	}
}

type toType struct {
	Field1 *int
	Field2 map[string]interface{}
}
