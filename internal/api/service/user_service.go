package service

import (
	"context"
	"film-management-api-golang/internal/api/repository"
	"film-management-api-golang/internal/dto"

	"gorm.io/gorm"
)

type (
	UserService interface {
		GetById(ctx context.Context, userId string, requesterId string) (dto.UserResponse, error)
	}

	userService struct {
		userRepository repository.UserRepository
		db             *gorm.DB
	}
)

func NewUser(userRepository repository.UserRepository,
	db *gorm.DB) UserService {
	return &userService{
		userRepository: userRepository,
		db:             db,
	}
}

func (s *userService) GetById(ctx context.Context, userId string, requesterId string) (dto.UserResponse, error) {
	user, err := s.userRepository.GetByIdWithFilmList(ctx, nil, userId)
	if err != nil {
		return dto.UserResponse{}, err
	}

	var filmLists []dto.FilmListResponse
	for _, filmlist := range user.FilmLists {
		// Visibility Logic: 
		// Hide private lists if the requester is not the owner
		if filmlist.Visibility == "private" && userId != requesterId {
			continue
		}

		filmLists = append(filmLists, dto.FilmListResponse{
			ID:         filmlist.ID.String(),
			FilmTitle:  filmlist.Film.Title,
			ListStatus: string(filmlist.ListStatus),
			Visibility: string(filmlist.Visibility),
		})
	}

	var reviewResponse []dto.ReviewResponse
	for _, review := range user.Reviews {
		reviewResponse = append(reviewResponse, dto.ReviewResponse{
			Film:    review.Film.Title,
			Rating:  review.Rating,
			Comment: review.Comment,
		})
	}

	return dto.UserResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Bio:         user.Bio,
		FilmLists:   filmLists,
		Reviews:     reviewResponse,
	}, nil
}
