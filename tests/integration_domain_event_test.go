package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tonybka/go-base-persistence/mock"
	"github.com/tonybka/go-base-persistence/model"
)

func TestAssignDomainEvent(t *testing.T) {
	testEventName := "Name"
	testEventID := "ID"

	events := []*mock.MockDomainEventStruct{{EventName: testEventName, EventID: testEventID}}

	model := model.BaseModel{}
	for _, event := range events {
		model.AddEvent(event)
	}

	assert.Greater(t, len(model.PendingEvents), 0)
	assert.Equal(t, testEventID, model.PendingEvents[0].ID())
	assert.Equal(t, testEventName, model.PendingEvents[0].Name())
}
