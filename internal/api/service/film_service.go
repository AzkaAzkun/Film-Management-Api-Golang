package service

import (
	"context"
	"film-management-api-golang/internal/api/repository"
	"film-management-api-golang/internal/dto"
	"film-management-api-golang/internal/entity"
	"film-management-api-golang/internal/utils"
	"fmt"
	"strings"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type (
	FilmService interface {
		Create(ctx context.Context, req dto.FilmCreateRequest) (dto.FilmCreateResponse, error)
	}

	filmService struct {
		filmRepository  repository.FilmRepository
		genreRepository repository.GenreRepository
		db              *gorm.DB
	}
)

func NewFilm(filmRepository repository.FilmRepository,
	genreRepository repository.GenreRepository,
	db *gorm.DB) FilmService {
	return &filmService{
		filmRepository:  filmRepository,
		genreRepository: genreRepository,
		db:              db,
	}
}

func (s *filmService) Create(ctx context.Context, req dto.FilmCreateRequest) (dto.FilmCreateResponse, error) {
	genreId := strings.Split(req.Genres, ",")

	genres, err := s.genreRepository.GetBatchById(ctx, nil, genreId)
	if err != nil {
		return dto.FilmCreateResponse{}, err
	}

	var creategenre []entity.FilmGenre
	for _, genre := range genres {
		creategenre = append(creategenre, entity.FilmGenre{
			GenreId: genre.ID,
		})
	}

	var createimage []entity.FilmImage
	for _, image := range req.Images {
		filename := fmt.Sprintf("film-%s-%s.%s", utils.ToSlug(req.Title), ulid.Make(), utils.GetExtensions(image.Filename))
		if err := utils.UploadFile(image, filename); err != nil {
			return dto.FilmCreateResponse{}, err
		}
		createimage = append(createimage, entity.FilmImage{
			ImagePath: filename,
		})
	}

	createResult, err := s.filmRepository.Create(ctx, nil, entity.Film{
		Title:         req.Title,
		Synopsis:      req.Synopsis,
		AiringStatus:  entity.AiringStatus(req.AiringStatus),
		TotalEpisodes: req.TotalEpisodes,
		ReleaseDate:   req.ReleaseDate,
		Images:        createimage,
		Genres:        creategenre,
	})
	if err != nil {
		return dto.FilmCreateResponse{}, err
	}

	return dto.FilmCreateResponse{
		ID: createResult.ID.String(),
	}, nil
}
