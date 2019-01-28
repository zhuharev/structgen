package models

import "time"

type TicketType int

const (
	TicketModerationRequestPermissions TicketType = iota + 1
	TicketOfferModeration
)

type Ticket struct {
	tableName struct{} `sql:"tickets,alias:t" pg:",discard_unknown_columns"`

	Type      TicketType `sql:"type" json:"type"`
	Title     string     `sql:"title" json:"title"`
	Text      string     `sql:"text" json:"text"`
	ID        int        `sql:"id" json:"id"`
	CreatedAt time.Time  `sql:"createdAt" json:"createdAt"`
	UpdatedAt time.Time  `sql:"updatedAt" json:"updatedAt"`
}
