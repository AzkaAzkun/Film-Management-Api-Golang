package repository

import (
	"context"
	"film-management-api-golang/internal/entity"

	"gorm.io/gorm"
)

type (
	FilmListRepository interface {
		Create(ctx context.Context, tx *gorm.DB, film entity.FilmList) (entity.FilmList, error)
	}

	filmListRepository struct {
		db *gorm.DB
	}
)

func NewFilmList(db *gorm.DB) FilmListRepository {
	return &filmListRepository{
		db: db,
	}
}

func (r *filmListRepository) Create(ctx context.Context, tx *gorm.DB, filmlist entity.FilmList) (entity.FilmList, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&filmlist).Error; err != nil {
		return entity.FilmList{}, err
	}

	return filmlist, nil
}
