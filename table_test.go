package tabular

import (
	"testing"
)

func TestTable(t *testing.T) {
	table := Table{
		fields: []Field{
			{Name: "name", Type: String},
		},
	}
	table.Append(map[string]string{"name": "Andrews"})
	table.Append(map[string]string{"name": "Tarsis"})

	if len(table.data) != 2 {
		t.Fatalf("table data length should be 2, got %q", len(table.data))
	}
}

func TestFields(t *testing.T) {
	table := Table{}
	table.SetFields([]string{"a", "b"})
	expected := []Field{
		{Name: "a", Type: String},
		{Name: "b", Type: String},
	}
	for i, f := range expected {
		if f.Name != table.fields[i].Name {
			t.Fatalf("field name is wrong. Expect %s, got %s", f.Name, table.fields[i].Name)
		}
		if f.Type != table.fields[i].Type {
			t.Fatalf("field type is wrong. Expect %s, got %s", f.Type, table.fields[i].Type)
		}
	}
}
