package entity

type AdminUser struct {
	BaseModel
	Name         string `json:"name"`
	Email        string `gorm:"uniqueIndex" json:"email"`
	PasswordHash string `json:"-"`
	IsActive     bool   `json:"is_active"`
}
