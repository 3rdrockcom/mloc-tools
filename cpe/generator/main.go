package main

import (
	"flag"
	"os"
)

var doSeed bool

func init() {
	flag.BoolVar(&doSeed, "seed", false, "seed the database")

	flag.Parse()
}

func main() {
	// Database
	db = NewDB("faker.db")
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
