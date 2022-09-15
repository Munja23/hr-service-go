package model

import (
	"time"
)

// TODO define model data types
type User struct {
	Id          *int       `json:"id,omitempty" db:"id"`
	Email       *string    `json:"email,omitempty" db:"email"`
	Password    *[32]byte  `json:"password,omitempty" db:"password"`
	FirstName   *string    `json:"first_name,omitempty" db:"first_name"`
	LastName    *string    `json:"last_name,omitempty" db:"last_name"`
	DateOfBirth *time.Time `json:"dateofbirth,omitempty" db:"dateofbirth"`
	Address     *string    `json:"address,omitempty" db:"address"`
	Contact     *string    `json:"phonenumber,omitempty" db:"phonenumber"`
	Degree      *string    `json:"collegedegree,omitempty" db:"collegedegree"`
	HireDate    *time.Time `json:"hiredate,omitempty" db:"hiredate"`
	IsAdmin     *bool      `json:"isadmin,omitempty" db:"isadmin"`
	PTO         *int       `json:"pto,omitempty" db:"pto"`
	CreateTime  *time.Time `json:"createtime,omitempty" db:"createtime"`
}
