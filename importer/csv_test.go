package importer

import (
	"strings"
	"testing"
)

func TestCSVImporter(t *testing.T) {
	data := `first_name,last_name,username
Andrews,Medina,andrewsmedina`
	csvImpoter := &CSV{}

	_, err := csvImpoter.Import(strings.NewReader(data))

	if err != nil {
		t.Fatalf("expect nil, got %q", err)
	}
}
