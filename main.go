package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(`
Usage: go-blitlink <dbfile> <command> [<args>]

Valid commands: insert, query, update, delete

insert
------
Arguments: <text> <link> <title> <shortcut>
Example: go-blitlink mydb.db insert "Hello World" "http://google.com" "Google" "g"

All fields are required, empty strings are permitted
Example: go-blitlink mydb.db insert "Hello World" "http://google.com" "" "g"

query
-----
Arguments: <text>
Example: go-blitlink mydb.db query "Hello World"

Performs a full-text search across all columns
Exact matches for the "shortcut" column are listed first

update
------
Arguments: <id> <text> <link> <title> <shortcut>
Example: go-blitlink mydb.db update 1 "Hello World" "http://github.com" "GitHub" "g"

All fields are required, empty strings are permitted
Example: go-blitlink mydb.db update 1 "Hello World" "http://github.com" "" "g"

delete
------
Arguments: <id>
Example: go-blitlink mydb.db delete 1
		`)
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
