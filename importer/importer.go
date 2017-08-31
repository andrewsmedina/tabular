package importer

import (
	"github.com/andrewsmedina/tabular"
)

var importBackends map[string]ImportBackend

// ImportBackend represets a data importer
type ImportBackend interface {
	Import(config map[string]string) (*tabular.Table, error)
}

// Register registers a new import backend
func Register(name string, backend ImportBackend) {
	if importBackends == nil {
		importBackends = map[string]ImportBackend{}
	}
	importBackends[name] = backend
}

// Get returns a import backend by name
func Get(name string) ImportBackend {
	return importBackends[name]
}
