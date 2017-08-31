package tabular

import (
	"reflect"
	"errors"
)

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
	fields []Field
	data   []interface{}
}

// SetFields add fields and respectives types in table
func (t *Table) SetFields(data interface{}) error {
	s := reflect.ValueOf(data)
    if s.Kind() != reflect.Slice {
        return errors.New("data given a non-slice type")
	}
	for i:=0; i<s.Len(); i++ {
		field := Field{
			Name: s.Index(i).String(),
			Type: String,
		}
		t.fields = append(t.fields, field)
	}
	return nil
}

// Append adds a row to a table data
func (t *Table) Append(row interface{}) {
	if t.data == nil {
		t.data = []interface{}{}
	}

	t.data = append(t.data, row)
}
