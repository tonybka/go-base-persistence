package account

import (
	"encoding/json"
	"strconv"
	"time"
)

type AccountCreatedEvent struct {
	id int
}

func (event *AccountCreatedEvent) Name() string {
	return "event.account.created"
}

func (event *AccountCreatedEvent) ToJson() (string, error) {
	bEvent, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(bEvent), nil
}

func (event *AccountCreatedEvent) ID() string {
	return strconv.Itoa(event.id)
}

func (event *AccountCreatedEvent) OccurredAt() time.Time {
	return time.Now()
}
