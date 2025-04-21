package repository

import (
	"context"
	"film-management-api-golang/internal/dto"
	"film-management-api-golang/internal/entity"
	"film-management-api-golang/internal/pkg/meta"

	"gorm.io/gorm"
)

type (
	FilmRepository interface {
		Create(ctx context.Context, tx *gorm.DB, film entity.Film) (entity.Film, error)
		GetById(ctx context.Context, tx *gorm.DB, filmId string) (entity.Film, error)
		GetAllPaginated(ctx context.Context, tx *gorm.DB, metareq meta.Meta) (dto.GetAllFilmPaginatedResponse, error)
	}

	filmRepository struct {
		db *gorm.DB
	}
)

func NewFilm(db *gorm.DB) FilmRepository {
	return &filmRepository{
		db: db,
	}
}

func (r *filmRepository) Create(ctx context.Context, tx *gorm.DB, film entity.Film) (entity.Film, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&film).Error; err != nil {
		return entity.Film{}, err
	}

	return film, nil
}

func (r *filmRepository) GetById(ctx context.Context, tx *gorm.DB, filmId string) (entity.Film, error) {
	if tx == nil {
		tx = r.db
	}

	var film entity.Film
	if err := tx.WithContext(ctx).Where("id = ?", filmId).Take(&film).Error; err != nil {
		return entity.Film{}, err
	}

	return film, nil
}

func (r *filmRepository) GetAllPaginated(ctx context.Context, tx *gorm.DB, metareq meta.Meta) (dto.GetAllFilmPaginatedResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var result []dto.GetAllFilmResponse
	query := tx.WithContext(ctx).Model(&entity.Film{})
	query = WithFilters(query, &metareq, AddModels(entity.Film{}))
	subQuery := r.db.
		Select("film_id, AVG(rating) as average_rating").
		Table("us_reviews").
		Group("film_id")
	query = query.
		Select("films.*, avg_ratings.average_rating").
		Joins("LEFT JOIN (?) as avg_ratings ON avg_ratings.film_id::uuid = films.id", subQuery).Scan(&result)

	return dto.GetAllFilmPaginatedResponse{
		Data: result,
		Meta: metareq,
	}, query.Error
}
