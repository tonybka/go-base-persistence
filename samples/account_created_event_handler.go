package account

import (
	"fmt"

	"github.com/tonybka/go-base-ddd/domain/event"
)

// AccountCreatedEventHandler triggered once new account created
type AccountCreatedEventHandler struct {
}

func (handler *AccountCreatedEventHandler) Notify(event event.IBaseDomainEvent) error {
	fmt.Println("AccountCreatedEventHandler.Notify: get notified")
	return nil
}
