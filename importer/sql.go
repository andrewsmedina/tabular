package importer

import (
	"database/sql"

	"github.com/andrewsmedina/tabular"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	Register("sql", &SQL{})
}

// SQL represents a sql importer
type SQL struct{}

// Import imports a sql data into a tabular.Table
func (c *SQL) Import(config map[string]string) (*tabular.Table, error) {
	table := &tabular.Table{}

	db, err := sql.Open(config["database"], config["uri"])
	if err != nil {
		return table, err
	}

	rows, err := db.Query(config["query"])
	if err != nil {
		return table, err
	}

	defer rows.Close()
	for rows.Next() {
		var row interface{}
		if err := rows.Scan(&row); err != nil {
			return table, err
		}

		table.Append(row)
	}
	return table, nil
}
