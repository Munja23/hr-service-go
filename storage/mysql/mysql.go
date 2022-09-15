package mysql

import (
	"gitea.delsystems.academy/academy/hr-service-go/conf"
	"github.com/mailru/dbr"
)

type MySQL struct {
	Connection *dbr.Connection
	Dbcfg      conf.DbCfg
}

const (
	OrganisationMySQLTableName = "tbl_hr_application_organisation"
)
