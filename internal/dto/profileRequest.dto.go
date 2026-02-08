package dto

import (
	"mime/multipart"
	"time"
)

type GetProfileRequest struct {
	ID        int       `json:"id"`
	AccountID int       `json:"account_id"`
	UserName  string    `json:"user_name"`
	UserImage string    `json:"user_image"`
	UserBio   string    `json:"user_bio"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ProfileResponse struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	UserImage string    `json:"user_image"`
	UserBio   string    `json:"user_bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InputProfileRequest struct {
	AccountID int                   `form:"account_id"`
	UserName  string                `form:"user_name"`
	UserImage *multipart.FileHeader `form:"user_image"`
	UserBio   string                `form:"user_bio"`
}
