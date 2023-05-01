package models

import "time"

type Session struct {
	sessionID    int
	sessionName  string
	creationDate time.Time

	Histories []History
}
