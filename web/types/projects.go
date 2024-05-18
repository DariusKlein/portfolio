package types

type Projects struct {
	projects []Project
}

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	ImageUrl    string `json:"image_url"`
}
