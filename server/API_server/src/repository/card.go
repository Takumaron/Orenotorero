package repository

import (
	"orenotorero/db/Model"
	"time"
)

type CardRepository interface {
	InsertCard(userId string, title string, kanbanId, position int) error
	UpdateCardTitle(id int, title string) error
	UpdateCardDeadLine(id int, deadline time.Time) error
	SelectAll(kanbanId int) ([]model.Card, error)
	DeleteCard(userId string, cardId int) error
}
