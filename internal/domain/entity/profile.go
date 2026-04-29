package entity

type Profile struct {
	BaseModel
	FullName    string `json:"full_name"`
	Headline    string `json:"headline"`
	SubHeadline string `json:"sub_headline"`
	Bio         string `json:"bio"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Location    string `json:"location"`
	AvatarPath  string `json:"avatar_path"`
	ResumePath  string `json:"resume_path"`
}
