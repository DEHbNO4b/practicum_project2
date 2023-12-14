package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var (
		storagePath     string
		migrationsPath  string
		migrationsTable string
	)
	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "migration table")
	flag.Parse()

	if storagePath == "" {
		panic("storage path is required")
	}
	if migrationsPath == "" {
		panic("migrations path is required")
	}
	m, err := migrate.New(
		"file://"+migrationsPath,
		"postgres://practicum:practicum@localhost:5432/practicum?sslmode=disable",
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
