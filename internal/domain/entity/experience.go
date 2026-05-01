package entity

type Experience struct {
	BaseModel
	Company     string `json:"company"`
	Position    string `json:"position"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	IsCurrent   bool   `json:"is_current"`
	IsActive    bool   `json:"is_active"`
	SortOrder   int    `json:"sort_order"`
}
