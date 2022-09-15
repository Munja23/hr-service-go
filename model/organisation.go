package model

import "time"

type Organisation struct {
	Id         *int       `json:"id,omitempty" db:"id"`
	Name       *string    `json:"name,omitempty" db:"name"`
	AdminEmail *string    `json:"email,omitempty" db:"email"`
	CreateTime *time.Time `json:"-" db:"create_time"`
}
