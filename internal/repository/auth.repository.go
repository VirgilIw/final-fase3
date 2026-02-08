package repository

import (
	"context"
	"log"

	"github.com/virgilIw/final-fase3/internal/dto"
	"github.com/virgilIw/final-fase3/internal/model"
)

type AuthRepo interface {
	Register(ctx context.Context, db DBTX, req dto.RegisterRequest) error
	Login(ctx context.Context, db DBTX, req dto.LoginRequest) (model.Account, error)
}

// Kenapa Exec()?
// karena INSERT tidak mengembalikan row.
type AuthRepository struct{}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (r *AuthRepository) Register(ctx context.Context, db DBTX, req dto.RegisterRequest) error {
	query := "INSERT INTO public.account (email,password) VALUES ($1,$2)"

	// perintah untuk jalankan query
	_, err := db.Exec(ctx, query, req.Email, req.Password)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (l *AuthRepository) Login(ctx context.Context, db DBTX, req dto.LoginRequest) (model.Account, error) {
	query := "SELECT id, email, password, role FROM public.account where email = $1;"

	row := db.QueryRow(ctx, query, req.Email)
	// perlu di scan karna nanti dapat data mentah
	var account model.Account
	err := row.Scan(
		&account.ID,
		&account.Email,
		&account.Password,
		&account.Role,
	)

	if err != nil {
		log.Println(err.Error())
		return model.Account{}, err
	}
	return account, nil
}
