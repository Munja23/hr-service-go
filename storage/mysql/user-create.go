package mysql

import (
	"gitea.delsystems.academy/academy/hr-service-go/model"
)

func (m MySQL) UserCreate(user model.User) (int, error) {
	session := m.Connection.NewSession(nil)

	builder := session.InsertInto("tbl_hr_application_user").Columns("email", "password", "creationdate", "isadmin")

	builder.Record(user)

	res, err := builder.Exec()

	if err != nil {
		return 0, err
	}

	i, err := res.RowsAffected()

	return int(i), err
}
