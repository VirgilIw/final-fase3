package model

import (
	"database/sql"
	"time"
)

type Account struct {
	ID          int          `db:"id"`
	Email       string       `db:"email"`
	Password    string       `db:"password"`
	Photo       string       `db:"photo"`
	Role        string       `db:"role"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
	LastLoginAt sql.NullTime `db:"lastlogin_at"`
}
