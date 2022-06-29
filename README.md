# go-blitlink

This project uses the go-sqlite3 package which must be built with CGO_ENABLED.
This project also uses the 'fts5' module for full-text search, which must be enabled with a build tag.

The included `build.sh` script will run this for you:

```
sh build.sh
CGO_ENABLED=1 go build --tags "fts5"
```

## Usage

```
go-blitlink [<dbfile>] [<command>] [<args>]
```

If the database file does not exist, it will be created.
If no command is specified, the program will read the number of entries in the database and exit.

Run the program without any arguments to see more detailed usage information.