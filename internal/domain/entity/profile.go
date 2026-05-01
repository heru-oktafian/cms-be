package entity

type Profile struct {
	BaseModel
	FullName         string `json:"full_name"`
	Headline         string `json:"headline"`
	SubHeadline      string `json:"sub_headline"`
	HeroDescription  string `json:"hero_description"`
	AboutDescription string `json:"about_description"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Location         string `json:"location"`
	AvatarPath       string `json:"avatar_path"`
	ResumePath       string `json:"resume_path"`
}
