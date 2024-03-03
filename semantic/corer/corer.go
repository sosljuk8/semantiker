package corer

import "semantiker/orm/ent"

// get all brandkeys
type BrandKey interface {
	GetCoresAds() []string
}

// get core by name of brandkey
type CoreStore interface {
	CheckCoreByName(name string) (*ent.Cores, error)
}

// if exists, continue, else create new core
