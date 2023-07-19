package db_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
	"github.com/joshuaetim/gallery_go/domain/model"
	"github.com/joshuaetim/gallery_go/domain/repository/db"
	"github.com/joshuaetim/gallery_go/factory"
	"github.com/stretchr/testify/assert"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type query map[string]interface{}

func TestUserSave(t *testing.T) {
	initTest(t)

	dbConn := db.DB()

	var user model.User
	user.Firstname = "Joshua Etim"
	user.Lastname = "Etim"
	user.Email = "jetimworks@gmail.com"
	user.Password = "password"

	ur := db.NewUserRepository(dbConn)

	u, err := ur.AddUser(user)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, u.ID)
	assert.EqualValues(t, u.Firstname, "Joshua Etim")
}

func TestUserDuplicateEmail(t *testing.T) {
	initTest(t)
	dbConn := db.DB()

	user1 := model.User{
		Firstname: "Josh",
		Lastname:  "Etim",
		Email:     "jetimworks@gmail.com",
		Password:  "password",
	}
	user2 := model.User{
		Firstname: "Josh",
		Lastname:  "Etim",
		Email:     "jetimworks@gmail.com",
		Password:  "password",
	}
	ur := db.NewUserRepository(dbConn)

	u1, err := ur.AddUser(user1)
	assert.Nil(t, err)
	assert.EqualValues(t, u1.Email, "jetimworks@gmail.com")

	u2, err := ur.AddUser(user2)
	assert.NotNil(t, err)
	assert.EqualValues(t, u2.ID, 0)
}

func TestUserGet(t *testing.T) {
	initTest(t)
	dbConn := db.DB()

	u1, err := factory.SeedUser(dbConn)
	assert.Nil(t, err)
	ur := db.NewUserRepository(dbConn)

	res, err := ur.GetMap(query{"id": u1.ID})
	u2 := res[0]

	assert.Nil(t, err)
	assert.EqualValues(t, u2.Email, u1.Email)
}

func TestUserGetByEmail(t *testing.T) {
	initTest(t)
	dbConn := db.DB()

	u1, err := factory.SeedUser(dbConn)
	assert.Nil(t, err)
	ur := db.NewUserRepository(dbConn)

	res, err := ur.GetMap(query{"email": u1.Email})
	u2 := res[0]

	assert.Nil(t, err)
	assert.EqualValues(t, u1.ID, u2.ID)
}

func TestUserGetAll(t *testing.T) {
	initTest(t)
	dbConn := db.DB()

	var users []model.User
	for i := 0; i < 4; i++ {
		u, err := factory.SeedUser(dbConn)
		assert.Nil(t, err)
		users = append(users, u)
	}
	fmt.Println(len(users))

	ur := db.NewUserRepository(dbConn)
	allUsers, err := ur.GetMap(query{"1": "1"})
	assert.Nil(t, err)
	assert.EqualValues(t, len(users), len(allUsers))
}

func TestUserUpdate(t *testing.T) {
	initTest(t)
	dbConn := db.DB()

	u, err := factory.SeedUser(dbConn)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, u.ID)

	u.Email = "changed@gmail.com"

	ur := db.NewUserRepository(dbConn)
	u2, err := ur.UpdateUser(u)
	assert.Nil(t, err)

	assert.EqualValues(t, "changed@gmail.com", u2.Email)
}

func TestUserDelete(t *testing.T) {
	initTest(t)
	dbConn := db.DB()

	u, err := factory.SeedUser(dbConn)
	assert.Nil(t, err)

	ur := db.NewUserRepository(dbConn)

	err = ur.DeleteUser(u)
	assert.Nil(t, err)

	_, err = ur.GetMap(query{"id": u.ID})
	assert.NotNil(t, err)
}

func initTest(t *testing.T) {
	err := godotenv.Load(basepath + "/../" + ".env_test")
	if err != nil {
		t.Fatal(err)
	}
}
