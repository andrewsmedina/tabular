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
