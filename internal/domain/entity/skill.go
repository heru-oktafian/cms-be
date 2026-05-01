package entity

type Skill struct {
	BaseModel
	Name      string `json:"name"`
	Category  string `json:"category"`
	Level     int    `json:"level"`
	IconPath  string `json:"icon_path"`
	IsActive  bool   `json:"is_active"`
	SortOrder int    `json:"sort_order"`
}
