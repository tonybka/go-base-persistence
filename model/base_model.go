package model

import (
	"errors"

	"github.com/tonybka/go-base-ddd/domain/event"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	PendingEvents []event.IBaseDomainEvent `gorm:"-"`
}

func NewBaseModel(id uint) BaseModel {
	model := gorm.Model{ID: uint(id)}
	return BaseModel{Model: model}
}

// publishOngoingEvents publish all events that occurred
func (model *BaseModel) publishOngoingEvents(tx *gorm.DB) error {
	publisher := GetDomainEventPublisher()
	if publisher == nil {
		return errors.New("invalid domain event publisher")
	}

	err := publisher.Publish(tx, model.PendingEvents...)
	if err != nil {
		return err
	}

	return nil
}

// AddEvent add event to ready for publishing.
// We are not allowed to convert a []T to an []interface{}, so we have to add the event one by one
func (model *BaseModel) AddEvent(domainEvent event.IBaseDomainEvent) {
	if len(model.PendingEvents) == 0 {
		model.PendingEvents = make([]event.IBaseDomainEvent, 0)
	}

	model.PendingEvents = append(model.PendingEvents, domainEvent)
}

func (model *BaseModel) AfterSave(tx *gorm.DB) (err error) {
	return model.publishOngoingEvents(tx)
}

func (model *BaseModel) AfterDelete(tx *gorm.DB) (err error) {
	return model.publishOngoingEvents(tx)
}
