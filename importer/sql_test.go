package importer

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSQLImport(t *testing.T) {
	db, err := sql.Open("sqlite3", "my.db")
	if err != nil {
		t.Fatalf("Error should be nil. Got %q", err)
	}

	_, err = db.Exec("CREATE TABLE userinfo (username VARCHAR(64) NULL);")
	if err != nil {
		t.Fatalf("Error should be nil. Got %q", err)
	}

	stmt, err := db.Prepare("INSERT INTO userinfo (username) values(?)")
	if err != nil {
		t.Fatalf("Error should be nil. Got %q", err)
	}

	_, err = stmt.Exec("andrews")
	if err != nil {
		t.Fatalf("Error should be nil. Got %q", err)
	}

	c := SQL{}
	_, err = c.Import(map[string]string{
		"database": "sqlite3",
		"uri":      "my.db",
		"query":    "SELECT username FROM userinfo",
	})
	if err != nil {
		t.Fatalf("Error importing sql. Got %s", err.Error())
	}
}
