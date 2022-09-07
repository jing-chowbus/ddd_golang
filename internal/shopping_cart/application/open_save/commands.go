package open_save

import (
	"ddd/internal/shopping_cart/domain"
	"time"

	"github.com/google/uuid"
)

type OpenCommand struct {
	Time        time.Time
	OrderSource string
	ServiceType string
	EmployeeID  uuid.UUID
}

type SaveCommand struct {
	Time         time.Time
	EmployeeID   uuid.UUID
	ShoppingCart domain.ShoppingCart
}
