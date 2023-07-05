package db

import (
	"fmt"

	"github.com/joshuaetim/akiraka3/domain/model"
	"github.com/joshuaetim/akiraka3/domain/repository/interfaces"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) AddUser(user model.User) (model.User, error) {
	return user, r.db.Create(&user).Error
}

func (r *userRepo) GetMap(query map[string]interface{}) ([]model.User, error) {
	var queryString string
	var fields []interface{}
	var user []model.User
	for k, v := range query {
		if queryString != "" {
			queryString = " " + queryString + " AND "
		}
		queryString = fmt.Sprintf("%s%s = ?", queryString, k)
		fields = append(fields, v)
	}
	// fields[0]
	var queryMain []interface{}
	queryMain = append(queryMain, queryString)
	queryMain = append(queryMain, fields...)

	return user, r.db.Find(&user, queryMain...).Error
}

func (r *userRepo) CountUsers() int {
	type Result struct {
		Total int
	}
	var result Result
	r.db.Raw("select count(*) as total from users").Scan(&result)

	return result.Total
}

func (r *userRepo) UpdateUser(user model.User) (model.User, error) {
	return user, r.db.Model(&user).Updates(&user).Error
}

func (r *userRepo) DeleteUser(user model.User) error {
	// exists?
	if err := r.db.First(&user).Error; err != nil {
		return err
	}
	return r.db.Delete(&user).Error
}
