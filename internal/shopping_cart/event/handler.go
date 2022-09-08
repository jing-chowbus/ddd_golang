package event

type NewShoppingCartEventHandler interface {
	Handle(NewEvent) error
}
