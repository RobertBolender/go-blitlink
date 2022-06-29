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
Usage: go-blitlink <dbfile> [<command>] [<args>]

Valid commands: insert, query, update, delete

If the database does not exist, it will be created.
When no command is specified, the program will print the number of entries in the database.

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

	if len(os.Args) < 3 {
		count(db)
		os.Exit(0)
	}

	switch os.Args[2] {
	case "insert":
		if len(os.Args) != 7 {
			log.Fatal("insert requires 4 arguments")
		}
		insert(db, os.Args[3], os.Args[4], os.Args[5], os.Args[6])
	case "update":
		if len(os.Args) != 8 {
			log.Fatal("update requires 5 arguments")
		}
		update(db, os.Args[3], os.Args[4], os.Args[5], os.Args[6], os.Args[7])
	case "delete":
		if len(os.Args) != 4 {
			log.Fatal("delete requires 1 argument")
		}
		delete(db, os.Args[3])
	case "query":
		log.Fatal("Not yet implemented: ", os.Args[2])
	default:
		log.Fatal("Unknown command: ", os.Args[2])
	}
}

func count(db *sql.DB) {
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

func insert(db *sql.DB, text, link, title, shortcut string) {
	stmt, err := db.Prepare("insert into blitlinks(text, link, title, shortcut) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(text, link, title, shortcut)
	if err != nil {
		log.Fatal(err)
	}
}

func update(db *sql.DB, id, text, link, title, shortcut string) {
	stmt, err := db.Prepare("update blitlinks set text = ?, link = ?, title = ?, shortcut = ? where rowid = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(text, link, title, shortcut, id)
	if err != nil {
		log.Fatal(err)
	}
}

func delete(db *sql.DB, id string) {
	stmt, err := db.Prepare("delete from blitlinks where rowid = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}
