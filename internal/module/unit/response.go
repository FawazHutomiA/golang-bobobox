package unit

import (
	"time"

	"github.com/google/uuid"
)

type UnitListResponse struct {
	ID          uuid.UUID `json:"ID"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type UnitDetailResponse struct {
	ID          uuid.UUID `json:"ID"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type UnitCreateResponse struct {
	ID          uuid.UUID `json:"ID"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type UnitUpdateResponse struct {
	ID          uuid.UUID `json:"ID"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	LastUpdated time.Time `json:"lastUpdated"`
}
