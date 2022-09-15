package api

import (
	"gitea.delsystems.academy/academy/hr-service-go/conf"
	"gitea.delsystems.academy/academy/hr-service-go/storage"
)

type API struct {
	// TODO make a map[string]storage.Storage
	Stg map[string]storage.Storage
}

func New(driver string, cfg conf.DbCfg) (*API, error) {
	// Create storage instance
	var stg storage.Storage
	stg, err := storage.StorageFactory(driver)
	if err != nil {
		return nil, err
	}

	res := new(API)
	res.Stg = make(map[string]storage.Storage)
	res.Stg[driver] = stg
	return res, nil
}
