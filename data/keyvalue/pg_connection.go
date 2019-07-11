package keyvalue

import (
	"fmt"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

type postgresDB struct {
	db *pg.DB
}

// NewpostgresDB returns a DBConnector interface backed by a postgres DB
func NewpostgresDB(addr, user, password, dbname string) DBConnector {

	options := &pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: dbname,
	}

	db := pg.Connect(options)

	return &postgresDB{
		db: db,
	}
}

// NewQuery returns an instance of a new QueryBuilder
func (d *postgresDB) NewQuery() QueryBuilder {
	return &pgSQLQuery{db: d.db}
}

// MigrateUp performs all the steps required to bring the backing DB into an initialised state
func (d *postgresDB) MigrateUp() error {
	return d.migrate("up")
}

// MigrateDown is the inverse of MigrateUp and intended to be used in testing environment to achieve a "clean slate".
func (d *postgresDB) MigrateDown() error {
	return d.migrate("reset")
}

func (d *postgresDB) migrate(cmd string) error {

	// Migrations
	migrations.DefaultCollection.DiscoverSQLMigrations("migrations/")
	_, _, _ = migrations.Run(d.db, "init")
	oldVersion, newVersion, err := migrations.Run(d.db, cmd)
	if err != nil {
		return err
	}
	if newVersion != oldVersion {
		fmt.Printf("Migration %v: from version %d to %d\n", cmd, oldVersion, newVersion)
	} else {
		fmt.Printf("Migration %v: not needed. version is %d\n", cmd, oldVersion)
	}

	return nil
}
