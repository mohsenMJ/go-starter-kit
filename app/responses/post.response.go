package responses

type PostResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}