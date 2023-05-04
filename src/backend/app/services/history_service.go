package services

import (
	"backend/app/gorms"
	"backend/app/models"
)

type HistoryService struct {
	g gorms.Gorms
}

func (h HistoryService) GetHistory(sessionId string) []models.History {
	var histories []models.History
	h.g.Gorm.Where("session_id = ?", sessionId).Order("creation_date asc").Find(&histories)

	return histories
}

func NewHistoryService(g gorms.Gorms) HistoryService {
	return HistoryService{g: g}
}
