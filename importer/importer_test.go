package importer

import (
	"testing"

	"github.com/andrewsmedina/tabular"
)

type FakeImporter struct{}

func (f *FakeImporter) Import(config map[string]string) (*tabular.Table, error) {
	return &tabular.Table{}, nil
}

func TestImporter(t *testing.T) {
	fake := &FakeImporter{}
	Register("fake", fake)

	result := Get("fake")

	if fake != result {
		t.Fatalf("expect %q, got %q", fake, result)
	}
}
