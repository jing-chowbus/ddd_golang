package checkout

import (
	"ddd/internal/shopping_cart/domain"
	"time"

	"github.com/google/uuid"
)

type CheckoutCommand struct {
	Time         time.Time
	EmployeeID   uuid.UUID
	ShoppingCart domain.ShoppingCart
}
