package dto

import (
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
)
