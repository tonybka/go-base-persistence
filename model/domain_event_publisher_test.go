package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tonybka/go-base-persistence/mock"
)

// TestRegisterEventHandler ensure handler registration works
func TestRegisterEventHandler(t *testing.T) {
	eventPublisher := InitDomainEventPublisher()
	assert.NotNil(t, eventPublisher)

	handler := &mock.MockEventHandler{}
	event := &mock.MockDomainEventStruct{}

	eventPublisher.RegisterSubscriber(event, handler)
}

// TestPublishDomainEvent ensure the handler was notified once event triggered
func TestPublishDomainEvent(t *testing.T) {
	eventPublisher := InitDomainEventPublisher()
	assert.NotNil(t, eventPublisher)

	handler := &mock.MockEventHandler{}
	event := &mock.MockDomainEventStruct{}

	eventPublisher.RegisterSubscriber(event, handler)

	// Before notification
	assert.False(t, handler.Notified)

	eventPublisher.Publish(nil, event)

	// After notification
	assert.True(t, handler.Notified)
}
