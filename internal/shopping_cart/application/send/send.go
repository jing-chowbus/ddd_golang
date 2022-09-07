package send

import (
	shared "ddd/internal/shared/entity"
)

type SendingShoppingCart interface {
	Send(SendCommand) ([]shared.LineItem, error)
}
