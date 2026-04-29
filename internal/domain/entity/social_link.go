package entity

type SocialLink struct {
	BaseModel
	Platform  string `json:"platform"`
	Label     string `json:"label"`
	URL       string `json:"url"`
	IconPath  string `json:"icon_path"`
	SortOrder int    `json:"sort_order"`
}
