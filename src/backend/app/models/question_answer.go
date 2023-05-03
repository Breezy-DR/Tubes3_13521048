package models

type QuestionNA struct {
	QuestionID int    `gorm:"primaryKey;auto_increment;not_null;column:question_id"`
	Question   string `gorm:"not_null;column:question"`
	Answer     string `gorm:"not_null;column:answer"`
}

func (QuestionNA) TableName() string {
	return "question_answer"
}
