package entity

import "github.com/google/uuid"

type Info struct {
	ID           uuid.UUID
	RestaurantID uuid.UUID
	Name         string
}
