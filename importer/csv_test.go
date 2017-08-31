package importer

import (
	"testing"
)

func TestCSVImport(t *testing.T) {
	c := CSV{}
	_, err := c.Import(map[string]string{"path": "testdata/test.csv"})
	if err != nil {
		t.Fatalf("Error importing csv. Got %s", err.Error())
	}
}
