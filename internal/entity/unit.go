package entity

import (
	"time"

	"github.com/google/uuid"
)

type Unit struct {
	ID          uuid.UUID `db:"id" json:"ID"`
	Name        string    `db:"name" json:"name"`
	Type        string    `db:"type" json:"type"`
	Status      string    `db:"status" json:"status"`
	LastUpdated time.Time `db:"lastUpdated" json:"lastUpdated"`
}

func (a *Unit) ToInsert() []interface{} {
	return []interface{}{
		a.ID,
		a.Name,
		a.Type,
		a.Status,
	}
}

func (a *Unit) ToUpdate() []interface{} {
	return []interface{}{
		a.ID,
		a.Status,
	}
}
