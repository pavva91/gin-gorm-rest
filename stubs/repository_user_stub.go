package stubs

import "github.com/pavva91/gin-gorm-rest/models"

type UserRepositoryStub struct {
	CreateUserFn    func(*models.User) (*models.User, error)
	ListUsersFn     func() ([]models.User, error)
	GetByIDFn       func(string) (*models.User, error)
	GetByEmailFn    func(string) (*models.User, error)
	GetByUsernameFn func(string) (*models.User, error)
	UpdateFn        func() (*models.User, error)
	DeleteFn        func() (*models.User, error)
}

func (stub UserRepositoryStub) CreateUser(user *models.User) (*models.User, error) {
	return stub.CreateUserFn(user)
}

func (stub UserRepositoryStub) ListUsers() ([]models.User, error) {
	return stub.ListUsersFn()
}

func (stub UserRepositoryStub) GetByID(id string) (*models.User, error) {
	return stub.GetByIDFn(id)
}

func (stub UserRepositoryStub) GetByEmail(email string) (*models.User, error) {
	return stub.GetByEmailFn(email)
}

func (stub UserRepositoryStub) GetByUsername(username string) (*models.User, error) {
	return stub.GetByUsernameFn(username)
}

func (stub UserRepositoryStub) Update(user *models.User) (*models.User, error) {
	return stub.UpdateFn()
}

func (stub UserRepositoryStub) Delete(id string) (*models.User, error) {
	return stub.DeleteFn()
}
