package importer

var importBackends map[string]ImportBackend

func init() {
	importBackends = map[string]ImportBackend{}
}

// ImportBackend represets a data importer
type ImportBackend interface {
	Import()
}

// Register registers a new import backend
func Register(name string, backend ImportBackend) {
	importBackends[name] = backend
}

// Get returns a import backend by name
func Get(name string) ImportBackend {
	return importBackends[name]
}
