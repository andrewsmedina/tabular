package tabular

// Type represents a field type
type Type string

const (
	// String represents a string/text field type
	String = Type("string")
)

// Field represents a field
type Field struct {
	Name string
	Type Type
}

// Table represents a table
type Table struct {
	Fields []Field
	Data   []interface{}
}

// Append adds a row to a table data
func (t *Table) Append(row interface{}) {
	if t.Data == nil {
		t.Data = []interface{}{}
	}

	t.Data = append(t.Data, row)
}
