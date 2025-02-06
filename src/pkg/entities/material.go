package entities

type Material struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ContentURL  string `json:"content_url"`
}
