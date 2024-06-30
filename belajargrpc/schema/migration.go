package schema

import (
	"database/sql"

	"github.com/GuiaBolso/darwin"
)

var migration = []darwin.Migration{
	{
		Version:     1,
		Description: "Add cities",
		Script: `
		CREATE TABLE cities (
			id    serial NOT NULL PRIMARY KEY,
			name  VARCHAR(15) NOT NULL UNIQUE
			
		);`,
	},
}

func Migrate(db *sql.DB) error {
	driver := darwin.NewGenericDriver(db, darwin.PostgresDialect{})

	d := darwin.New(driver, migration, nil)

	return d.Migrate()
}
