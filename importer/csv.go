package importer

import (
	"io/ioutil"
	".."
	"encoding/csv"
	"bytes"
)

// CSV represents a csv importer
type CSV struct{}

// Import imports a csv into a tabular.Table
func (c *CSV) Import(config map[string]string) (*tabular.Table, error) {
	table := &tabular.Table{}
	f, err := ioutil.ReadFile(config["path"])
	if err != nil {
		return table, err
	}
	rawData := csv.NewReader(bytes.NewReader(f))
	data, err := rawData.ReadAll()
	if err != nil {
		return table, err
	}
	err = table.SetFields(data[0])
	if err != nil {
		return table, err
	}
	for _, row := range data[1:] {
		table.Append(row)
	}
	return table, nil
}
