package models

import "time"

type History struct {
	historyID    int
	sessionID    int
	userEntry    string
	answer       string
	creationDate time.Time
}
