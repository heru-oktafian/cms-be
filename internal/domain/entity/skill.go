package entity

type Skill struct {
	BaseModel
	Name      string `json:"name"`
	Level     string `json:"level"`
	IconPath  string `json:"icon_path"`
	SortOrder int    `json:"sort_order"`
}
