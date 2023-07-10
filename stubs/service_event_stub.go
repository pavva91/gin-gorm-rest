package stubs

import "github.com/pavva91/gin-gorm-rest/models"

type EventServiceStub struct {
	ListEventsFn  func() ([]models.Event, error)
	CreateEventFn func() (*models.Event, error)
	GetByIdFn     func() (*models.Event, error)
	DeleteByIdFn  func() (*models.Event, error)
	SaveEventFn   func() (*models.Event, error)
}

func (stub EventServiceStub) ListAllEvents() ([]models.Event, error) {
	return stub.ListEventsFn()
}

func (stub EventServiceStub) CreateEvent(event *models.Event) (*models.Event, error) {
	return stub.CreateEventFn()
}

func (stub EventServiceStub) GetById(id string) (*models.Event, error) {
	return stub.GetByIdFn()
}

func (stub EventServiceStub) DeleteById(id string) (*models.Event, error) {
	return stub.DeleteByIdFn()
}

func (stub EventServiceStub) SaveEvent(event *models.Event) (*models.Event, error) {
	return stub.SaveEventFn()
}


