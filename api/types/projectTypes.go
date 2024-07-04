package types

type Project struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	DocURL      string `json:"doc_url,omitempty"`
}
