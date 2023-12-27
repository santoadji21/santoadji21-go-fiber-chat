package models

type Message struct {
	Model
	Username string `gorm:"not null" json:"username"`
	Message  string `gorm:"type:text;not null" json:"message"`
	Room     string `gorm:"not null" json:"room"`
}
