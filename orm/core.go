package orm

import (
	"semantiker/orm/ent"
)

// statuses
// 1. new
// 2. done
// 3. stop (if stopped by user
// 2. ready (for processing) if status is not done or stop and file exists

type CoreStore struct {
	db *ent.Client
}

func NewCoreStore() *CoreStore {
	return &CoreStore{
		db: GetConnection(),
	}
}

// if exists, continue, else create new core
func (g *CoreStore) CheckCoreByName(name string) (*ent.Cores, error) {
	g.db.Close()

	return nil, nil
}

// get first ready core
func (g *CoreStore) GetReadyCore() (*ent.Cores, error) {
	g.db.Close()

	return nil, nil
}
