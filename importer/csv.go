package importer

import (
	"io/ioutil"
	"fmt"
	".."
	"strings"
	"csv"
)

// CSV represents a csv importer
type CSV struct{}

// Import imports a csv into a tabular.Table
func (c *CSV) Import(config map[string]string) (*tabular.Table, error) {
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
