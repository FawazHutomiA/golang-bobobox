package unit

import (
	"time"

	"github.com/google/uuid"
)

type UnitCreateRequest struct {
	ID          uuid.UUID `json:"ID"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type UnitUpdateRequest struct {
	ID          uuid.UUID `json:"ID"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	LastUpdated time.Time `json:"lastUpdated"`
}
