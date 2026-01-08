package dto

type CreatePostRequest struct {
	Caption string `json:"caption" binding:"max=500"`
}
