package entity

import "github.com/google/uuid"

type Preference struct {
	ID           uuid.UUID
	RestaurantID uuid.UUID
	QROrdering   bool
}
