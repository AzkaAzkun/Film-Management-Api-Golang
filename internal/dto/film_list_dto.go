package dto

type (
	FilmListRequest struct {
		FilmId     string `json:"film_id" binding:"required,uuid"`
		ListStatus string `json:"list_status" binding:"required"`
	}

	FilmListResponse struct {
		FilmTitle  string `json:"film_title"`
		ListStatus string `json:"list_status"`
	}
)
