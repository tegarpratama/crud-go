package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID 			uuid.UUID
	Name 		string
	Description string
    Category_id uuid.UUID
	CreatedAt 	time.Time
}