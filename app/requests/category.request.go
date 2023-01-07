package requests

type CategoryCreateRequest struct {
	Title string `json:"title" binding:"required"`
}

type CategoryUpdateRequest struct {
	Title string `json:"title" binding:"required"`
}
