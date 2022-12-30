package inputs

type PostCreateInput struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type PostUpdateInput struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}