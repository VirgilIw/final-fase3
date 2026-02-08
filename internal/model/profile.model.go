package model

import "time"

type Profile struct {
	ID        int       `db:"id"`
	AccountID int       `db:"account_id"`
	UserName  string    `db:"user_name"`
	UserImage string    `db:"user_image"`
	UserBio   string    `db:"user_bio"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
