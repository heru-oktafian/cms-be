package entity

type SocialLink struct {
	BaseModel
	Platform  string `json:"platform"`
	Label     string `json:"label"`
	URL       string `json:"url"`
	IconPath  string `json:"icon_path"`
	IsActive  bool   `json:"is_active"`
	SortOrder int    `json:"sort_order"`
}
