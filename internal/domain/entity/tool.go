package entity

type Tool struct {
	BaseModel
	Name      string `gorm:"size:100;not null" json:"name"`
	IconText  string `gorm:"size:50" json:"icon_text"`
	URL       string `gorm:"size:255" json:"url"`
	SortOrder int    `gorm:"default:0" json:"sort_order"`
	IsActive  bool   `gorm:"default:true" json:"is_active"`
}
