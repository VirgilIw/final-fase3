package repository

import (
	"context"

	"github.com/virgilIw/final-fase3/internal/dto"
	"github.com/virgilIw/final-fase3/internal/model"
)

type ProfileRepo interface {
	GetProfile(ctx context.Context, req dto.GetProfileRequest) (model.Profile, error)
}

type ProfileRepository struct{}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}

func (p *ProfileRepository) GetProfile(ctx context.Context, db DBTX, req dto.GetProfileRequest) (model.Profile, error) {
	query := `SELECT id,
       account_id,
       user_name,
       user_image,
       user_bio,
       created_at,
       updated_at
FROM public.users
WHERE id = $1;
`
	row := db.QueryRow(ctx, query, req.AccountID)
	var profile model.Profile

	err := row.Scan(
		&profile.ID,
		&profile.AccountID,
		&profile.UserName,
		&profile.UserImage,
		&profile.UserBio,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	if err != nil {
		return model.Profile{}, err
	}

	return profile, nil
}

func (r *ProfileRepository) InputProfile(
	ctx context.Context,
	db DBTX,
	req dto.InputProfileRequest,
	imagePath string,
) error {

	query := `
		UPDATE public.users
		SET
			user_name = $1,
			user_image = $2,
			user_bio = $3,
			updated_at = NOW()
		WHERE account_id = $4;
	`

	_, err := db.Exec(
		ctx,
		query,
		req.UserName,
		imagePath,
		req.UserBio,
		req.AccountID,
	)

	return err
}

// func (u *UserRepository) EditProfile(ctx context.Context, profileImg string, id int) (pgconn.CommandTag, error) {
// 	sql := "UPDATE public.users SET user_image = $2 WHERE id = $1"
// 	values := []any{id, profileImg}
// 	return u.db.Exec(ctx, sql, values...)
// }
