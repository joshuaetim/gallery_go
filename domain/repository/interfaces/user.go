package interfaces

import "github.com/joshuaetim/akiraka3/domain/model"

type UserRepository interface {
	AddUser(model.User) (model.User, error)
	GetMap(map[string]interface{}) ([]model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(model.User) error
	CountUsers() int
}
