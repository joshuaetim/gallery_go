package route

import (
	"github.com/joshuaetim/akiraka3/domain/repository/db"
	"github.com/joshuaetim/akiraka3/domain/repository/interfaces"
	"gorm.io/gorm"
)

func InitRepos(dbConn *gorm.DB) *interfaces.Repositories {
	return &interfaces.Repositories{
		User:  db.NewUserRepository(dbConn),
		Photo: db.NewPhotoRepository(dbConn),
		Like:  db.NewLikeRepository(dbConn),
	}
}
