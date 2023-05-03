package models

import "time"

type Session struct {
	SessionID    string    `gorm:"primaryKey;not_null;column:session_id"`
	SessionName  string    `gorm:"column:session_name"`
	CreationDate time.Time `gorm:"not_null;column:creation_date"`

	//Histories []History
}

func (Session) TableName() string {
	return "session"
}
