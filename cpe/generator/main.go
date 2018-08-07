package main

import (
	"flag"
	"os"
)

var dsn string
var doSeed bool

func init() {
	flag.StringVar(&dsn, "dsn", "", "e.g. user:password@localhost:3306/dbname?parseTime=true")
	flag.BoolVar(&doSeed, "seed", false, "seed the database")
	flag.Parse()
}

func main() {
	// Database
	db = NewDB(dsn)
	defer db.Close()

	if doSeed {
		// Migrate the schema
		DoMigrations()

		// Seed database with fake content
		s := NewSeeder()
		s.Run()

		os.Exit(0)
	}

	// Router
	r := NewRouter()
	r.Run()
}
