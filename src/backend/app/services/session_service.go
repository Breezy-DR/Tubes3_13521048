package services

import (
	"backend/app/gorms"
	"backend/app/models"
	util2 "backend/app/util"
	"fmt"
	"time"
)

type SessionService struct {
	g gorms.Gorms
}

func (s SessionService) GetNewSession(sessionName string) string {
	sessionId := util2.GenerateSessionId()
	s.g.Gorm.Create(&models.Session{SessionID: sessionId, SessionName: sessionName, CreationDate: time.Now()})

	return sessionId
}

func (s SessionService) GetNewestSessions() []models.Session {
	var sessions []models.Session
	s.g.Gorm.Order("creation_date desc").Limit(10).Find(&sessions)
	fmt.Println(sessions)
	return sessions
}

func (s SessionService) SessionExists(sessionId string) bool {
	var session = models.Session{SessionID: sessionId}
	res := s.g.Gorm.First(&session)

	return res.RowsAffected > 0
}

func NewSessionService(g gorms.Gorms) SessionService {
	return SessionService{g: g}
}
