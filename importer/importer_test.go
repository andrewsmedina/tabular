package importer

import (
	"testing"
)

type FakeImporter struct{}

func (f *FakeImporter) Import(config map[string]string) {}
func TestImporter(t *testing.T) {
	fake := &FakeImporter{}
	Register("fake", fake)

	result := Get("fake")

	if fake != result {
		t.Fatalf("expect %q, got %q", fake, result)
	}
}
