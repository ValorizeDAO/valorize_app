package main

import (
	gm "github.com/ShkrutDenis/go-migrations"
	gmStore "github.com/ShkrutDenis/go-migrations/store"
	"valorize.backend/db/migrations/list"
)

func main() {
	gm.Run(getMigrationsList())
}

func getMigrationsList() []gmStore.Migratable {
	return []gmStore.Migratable{
		&list.CreateUserTable{},
		&list.CreateContractsTable{},
	}
}
