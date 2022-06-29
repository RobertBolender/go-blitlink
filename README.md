# go-blitlink

A simple CLI tool for storing links for future use.  
Intended for use with [Raycast](https://www.raycast.com/) (MacOS) as a [community extension](https://github.com/RobertBolender/raycast-blitlink), although it can be used on any platform as a standalone binary.

## Features

* Create, update, and delete links
* Search for links with full-text search

## Building

This project uses the `mattn/go-sqlite3` package which must be built with `CGO_ENABLED`.
This project also uses the `'fts5'` sqlite module for full-text search, which must be enabled with a build tag.

The included `./script/build` script will run this for you:

```
./script/build
CGO_ENABLED=1 go build --tags "fts5"
```

## Usage

```
go-blitlink [<dbfile>] [<command>] [<args>]
```

If the database file does not exist, it will be created.
If no command is specified, the program will read the number of entries in the database and exit.

Run the program without any arguments to see more detailed usage information.

## Possible Future Changes

- Add the ability to sort the results by date_created or date_updated
- Add the ability to track how often a link is used for weighting "favorite" search results
- Change the command line arguments to use `--flags` instead of positional arguments
