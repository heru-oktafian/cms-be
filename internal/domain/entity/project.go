package entity

type Project struct {
	BaseModel
	Title         string `json:"title"`
	Slug          string `gorm:"uniqueIndex" json:"slug"`
	Summary       string `json:"summary"`
	Description   string `json:"description"`
	ThumbnailPath string `json:"thumbnail_path"`
	ProjectURL    string `json:"project_url"`
	RepoURL       string `json:"repo_url"`
	IsFeatured    bool   `json:"is_featured"`
	SortOrder     int    `json:"sort_order"`
}
