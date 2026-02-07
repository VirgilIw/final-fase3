package service

import (
	"context"
	"errors"
	"regexp"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/virgilIw/final-fase3/internal/dto"
	"github.com/virgilIw/final-fase3/internal/repository"
	"github.com/virgilIw/final-fase3/pkg/hash"
)

type AuthService struct {
	authRepository *repository.AuthRepository
	redis          *redis.Client
	db             *pgxpool.Pool
	hashConfig     *hash.HashConfig
}

func NewAuthService(authRepository *repository.AuthRepository, rdb *redis.Client, db *pgxpool.Pool, hashConfig *hash.HashConfig) *AuthService {
	return &AuthService{authRepository: authRepository, redis: rdb, db: db, hashConfig: hashConfig}
}

func (as *AuthService) Register(ctx context.Context, req dto.RegisterRequest) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, err := regexp.MatchString(emailRegex, req.Email)
	if err != nil {
		return err
	}

	if !matched {
		return errors.New("invalid email format")
	}

	// validasi password
	if len(req.Password) < 8 {
		return errors.New("password must be at least 6 characters")
	}

	hashedPwd, err := as.hashConfig.GenHash(req.Password)
	if err != nil {
		return err
	}

	req.Password = hashedPwd

	return as.authRepository.Register(ctx, as.db, req)
}

// LOGIN (password dibandingkan dengan hash)
func (as *AuthService) Login(ctx context.Context, req dto.LoginRequest) error {

	account, err := as.authRepository.Login(ctx, as.db, req)

	if err != nil {
		return err
	}

	isMatch, err := as.hashConfig.ComparePwdAndHash(req.Password, account.Password)
	if err != nil {
		return err
	}

	if !isMatch {
		return errors.New("invalid email or password")
	}
	return nil
}
