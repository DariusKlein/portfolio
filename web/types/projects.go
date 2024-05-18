package types

type Projects struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	ImageUrl    string `json:"image_url"`
	DocUrl      string `json:"doc_url"`
}
