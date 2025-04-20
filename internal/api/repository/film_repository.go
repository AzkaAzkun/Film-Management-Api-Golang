package repository

import (
	"context"
	"film-management-api-golang/internal/entity"

	"gorm.io/gorm"
)

type (
	FilmRepository interface {
		Create(ctx context.Context, tx *gorm.DB, film entity.Film) (entity.Film, error)
		GetById(ctx context.Context, tx *gorm.DB, filmId string) (entity.Film, error)
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
