package entity

type Skill struct {
	BaseModel
	Name      string `json:"name"`
	Category  string `json:"category"`
	Level     int    `json:"level"`
	IsActive  bool   `json:"is_active"`
	SortOrder int    `json:"sort_order"`
}
