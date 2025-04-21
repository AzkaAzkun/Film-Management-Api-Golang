package dto

import (
	"film-management-api-golang/internal/pkg/meta"
	"mime/multipart"
	"time"
)

type (
	FilmCreateRequest struct {
		Title         string                  `form:"title"`
		Synopsis      string                  `form:"synopsis"`
		AiringStatus  string                  `form:"airing_status"`
		TotalEpisodes int                     `form:"total_episodes"`
		ReleaseDate   time.Time               `form:"release_date" time_format:"2006-01-02 15:04:05"`
		Genres        string                  `form:"genres"`
		Images        []*multipart.FileHeader `form:"images"`
	}

	FilmCreateResponse struct {
		ID string `json:"id"`
	}

	GetAllFilmResponse struct {
		Title         string  `json:"title"`
		AiringStatus  string  `json:"airing_status"`
		TotalEpisodes int     `json:"total_episodes"`
		ReleaseDate   string  `json:"release_date"`
		AverageRating float32 `json:"average_rating"`
	}

	GetAllFilmPaginatedResponse struct {
		Data []GetAllFilmResponse
		Meta meta.Meta
	}
)
