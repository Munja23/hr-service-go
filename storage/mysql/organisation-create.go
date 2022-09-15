package mysql

import (
	"crypto/sha256"
	"time"

	"gitea.delsystems.academy/academy/hr-service-go/model"
)

const (
	DEFAULT_PASSWORD_LENGTH int = 16
)

func (m MySQL) OrganisationCreate(o model.Organisation) (model.Login, error) {
	pass := model.GeneratePassword(16)

	hash := sha256.Sum256([]byte(pass))

	now := time.Now()

	admin := model.User{
		Email:      o.AdminEmail,
		Password:   &hash,
		CreateTime: &now,
	}

	o.CreateTime = &now

	_, err := m.UserCreate(admin)
	if err != nil {
		return model.Login{}, err
	}

	sess := m.Connection.NewSession(nil)

	builder := sess.InsertInto(OrganisationMySQLTableName).Record(o)

	_, err = builder.Exec()
	if err != nil {
		return model.Login{}, err
	}

	return admin, nil
}
