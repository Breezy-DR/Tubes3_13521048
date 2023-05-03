package models

import "time"

type History struct {
	HistoryID    int       `gorm:"primaryKey;auto_increment;not_null;column:history_id"`
	SessionID    string    `gorm:"not_null;column:session_id"`
	UserEntry    string    `gorm:"not_null;column:user_entry"`
	Answer       string    `gorm:"not_null;column:answer"`
	CreationDate time.Time `gorm:"not_null;column:creation_date"`
}

func (History) TableName() string {
	return "history"
}
