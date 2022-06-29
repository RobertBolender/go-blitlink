package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go-blitlink <dbfile>")
	}

	db, err := sql.Open("sqlite3", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	setup := `
	create virtual table if not exists blitlinks using fts5(text, link, title, shortcut);
	`
	_, err = db.Exec(setup)
	if err != nil {
		log.Fatalf("%q: %s\n", err, setup)
	}

	countStmt := `
	select count(*) from blitlinks;
	`

	rows, err := db.Query(countStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, countStmt)
	}
	defer rows.Close()

	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Link journal entries: %d", count)
	}
}
