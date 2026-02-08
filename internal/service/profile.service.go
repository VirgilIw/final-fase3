package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/virgilIw/final-fase3/internal/dto"
	"github.com/virgilIw/final-fase3/internal/repository"
)

type ProfileService struct {
	profileRepo *repository.ProfileRepository
	redis       *redis.Client
	db          *pgxpool.Pool
}

func NewProfileService(profileRepo *repository.ProfileRepository, redis *redis.Client, db *pgxpool.Pool) *ProfileService {
	return &ProfileService{
		profileRepo: profileRepo,
		redis:       redis,
		db:          db,
	}
}

func (g *ProfileService) GetProfile(ctx context.Context, req dto.GetProfileRequest) (dto.ProfileResponse, error) {
	profile1, err := g.profileRepo.GetProfile(ctx, g.db, req)
	if err != nil {
		return dto.ProfileResponse{}, err
	}
	answer := dto.ProfileResponse{
		ID:        profile1.ID,
		UserName:  profile1.UserName,
		UserImage: profile1.UserImage,
		UserBio:   profile1.UserBio,
		// CreatedAt: profile1.CreatedAt,
		// UpdatedAt: profile1.UpdatedAt,
	}

	return answer, nil
}

func (s *ProfileService) InputProfile(
	ctx context.Context,
	req dto.InputProfileRequest,
	imagePath string,
) error {

	return s.profileRepo.InputProfile(
		ctx,
		s.db,
		req,
		imagePath,
	)
}
