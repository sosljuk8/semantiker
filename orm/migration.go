package orm

import (
	"context"
	"log"
	"semantiker/orm/ent"
)

type MigrationStore struct {
	db *ent.Client
}

func NewMigrationStore() *MigrationStore {
	return &MigrationStore{
		db: GetConnection(),
	}
}

func (m *MigrationStore) MigrateDb() {
	// Run the auto migration tool.
	if err := m.db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	defer m.db.Close()
}
