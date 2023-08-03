package model

type UserScore struct {
	ID      int    `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"size:15;not null;" json:"username"`
	Country string `gorm:"size:255;not null;" json:"-"`
	Score   int    `gorm:"default:0;" json:"score"`
}
