package stubs

import "github.com/pavva91/gin-gorm-rest/models"

type UserServiceStub struct {
	CreateFn        func() (*models.User, error)
	ListFn          func() ([]models.User, error)
	GetByIDFn       func() (*models.User, error)
	GetByEmailFn    func() (*models.User, error)
	GetByUsernameFn func() (*models.User, error)
	UpdateFn        func() (*models.User, error)
	DeleteFn        func() (*models.User, error)
}

func (stub UserServiceStub) Create(*models.User) (*models.User, error) {
	return stub.CreateFn()
}

func (stub UserServiceStub) List() ([]models.User, error) {
	return stub.ListFn()
}

func (stub UserServiceStub) GetByID(id string) (*models.User, error) {
	return stub.GetByIDFn()
}

func (stub UserServiceStub) GetByEmail(username string) (*models.User, error) {
	return stub.GetByEmailFn()
}

func (stub UserServiceStub) GetByUsername(username string) (*models.User, error) {
	return stub.GetByUsernameFn()
}

func (stub UserServiceStub) Update(user *models.User) (*models.User, error) {
	return stub.UpdateFn()
}

func (stub UserServiceStub) Delete(id string) (*models.User, error) {
	return stub.DeleteFn()
}
