package cmd

import "semantiker/orm"

type MigrationStore interface {
	MigrateDb()
}

type Migrator struct {
	mig MigrationStore
}

func (m Migrator) DoMigration() {
	m.mig.MigrateDb()

}

func MigrateSql() {
	// declare the migrator
	var md Migrator
	// set the migrator
	md.mig = orm.NewMigrationStore()
	// call the migrator
	md.DoMigration()
}
