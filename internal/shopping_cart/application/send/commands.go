package send

import (
	shared "ddd/internal/shared/entity"
	"time"

	"github.com/google/uuid"
)

type SendCommand struct {
	Time           time.Time
	EmployeeID     uuid.UUID
	ShoppingCartID uuid.UUID
	Items          []shared.LineItem
}
