package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

func main() {
	var (
		storagePath     string
		migrationsPath  string
		migrationsTable string
	)
	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "", "migration table")
	flag.Parse()

	if storagePath == "" {
		panic("storage path is required")
	}
	if migrationsPath == "" {
		panic("migrations path is required")
	}
	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("postgres://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migration change")

			return
		}
		panic(err)
	}
	fmt.Println("migrations applyed")

}
