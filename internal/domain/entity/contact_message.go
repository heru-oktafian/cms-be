package entity

import "time"

type ContactMessage struct {
	BaseModel
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	Phone         string     `json:"phone"`
	Subject       string     `json:"subject"`
	Message       string     `json:"message"`
	Status        string     `json:"status"`
	FollowUpNotes string     `json:"follow_up_notes"`
	FollowedUpAt  *time.Time `json:"followed_up_at"`
}
