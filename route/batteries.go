package route

import (
	"github.com/joshuaetim/akiraka3/domain/repository/db"
	"github.com/joshuaetim/akiraka3/domain/repository/interfaces"
	"gorm.io/gorm"
)

type repos struct {
	user  interfaces.UserRepository
	photo interfaces.PhotoRepository
}

func InitRepos(dbConn *gorm.DB) *repos {
	return &repos{
		user:  db.NewUserRepository(dbConn),
		photo: db.NewPhotoRepository(dbConn),
	}
}
