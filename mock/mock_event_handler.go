package mock

import "github.com/tonybka/go-base-ddd/domain/event"

type MockEventHandler struct {
	Notified bool
}

func (handler *MockEventHandler) Notify(event event.IBaseDomainEvent) error {
	handler.Notified = true
	return nil
}
