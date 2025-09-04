package unit

import (
	"github.com/google/uuid"
)

type UnitCreateRequest struct {
	ID     uuid.UUID `json:"ID"`
	Name   string    `json:"name"`
	Type   string    `json:"type"`
	Status string    `json:"status"`
}

type UnitUpdateRequest struct {
	ID     uuid.UUID `json:"ID"`
	Status string    `json:"status"`
}
