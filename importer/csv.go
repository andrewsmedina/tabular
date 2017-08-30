package importer

import (
	"encoding/csv"
	"io"

	"github.com/andrewsmedina/tabular"
)

// CSV represents a csv importer
type CSV struct{}

// Import imports a csv into a tabular.Table
func (c *CSV) Import(reader io.Reader) (*tabular.Table, error) {
	r := csv.NewReader(reader)

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	table := &tabular.Table{}

	for _, row := range records {
		table.Append(row)
	}
	return table, nil
}
