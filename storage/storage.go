package storage

import (
	"fmt"

	"gitea.delsystems.academy/academy/hr-service-go/model"
)

// TODO define storage interface
// And simple factory implementation
type Storage interface {
	UserRead(model.Queries) ([]model.User, error)
	UserCreate(model.User) (int, error)
	UserUpdate(model.Queries, model.User) (int, error)
	UserDelete(model.Queries) (int, error)

	OrganisationRead(model.Queries) ([]model.Organisation, error)
	OrganisationCreate(model.Organisation) (model.Login, error)
	OrganisationUpdate(model.Queries, model.Organisation) (int, error)
	OrganisationDelete(model.Queries) (int, error)
}

func StorageFactory(driver string) (Storage, error) {
	switch driver {
	default:
		return nil, fmt.Errorf("storageFactory error: Invalid or unsupported driver")
	}
}
