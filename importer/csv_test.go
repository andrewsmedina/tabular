package importer

import (
	"testing"
)

func TestCSVImport(t *testing.T) {
	c := CSV{}
	table, err := c.Import(map[string]string{"path": "testdata/test.csv"})
	if err != nil {
		t.Fatalf("Error importing csv. Got %s", err.Error())
	}
	if len(table.Fields) != 3 {
		t.Fatalf("Error importing csv. Want %d fields, got %d", 3, len(table.Fields))
	}
}
