package dto

type (
	ReactionRequest struct {
		ReviewId string `json:"review_id" binding:"required,uuid"`
		Status   string `json:"status" binding:"required"`
	}

	ReactionUpdate struct {
		Status string `json:"status" binding:"required"`
	}

	ReactionResponse struct {
		ID       string `json:"id"`
		ReviewId string `json:"review_id"`
		UserId   string `json:"user_id"`
		Status   string `json:"status"`
	}
)
