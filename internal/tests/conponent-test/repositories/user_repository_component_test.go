package test

import (
	"context"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
	"testing"
)

type UserRepositoryComponentTest struct {
	repository repositories.UserRepository
}

func NewUserRepositoryComponentTest(userRepository repositories.UserRepository) *UserRepositoryComponentTest {
	return &UserRepositoryComponentTest{
		repository: userRepository,
	}
}

var (
	UserId1   = "firstUser"
	Password1 = "password"
	UserId2   = "secondUser"
	Password2 = "password"
	UserId3   = "thirdUser"
	Password3 = "password"
)

var userRepositoryComponentTest *UserRepositoryComponentTest

func init() {
	logrus.Infof("Start Initializing User Repository Test Successfully")
	db, err := gorm.Open(sqlite.Open("../../../../data.db"), &gorm.Config{})
	if err != nil {
		logrus.Errorf("Open Gorm Mock error: %s", err.Error())
		return
	}

	userRepository := repositories.NewUserRepository(db)
	userRepositoryComponentTest = NewUserRepositoryComponentTest(userRepository)
	logrus.Infof("End Initializing User Repository Test Successfully")
}

func TestGetAllUser(t *testing.T) {
	users, err := userRepositoryComponentTest.repository.GetAll(context.TODO())

	assert.Equal(t, err, nil)
	assert.Equal(t, len(users), 3)

	assert.Equal(t, users[0].ID, UserId1)
	assert.Equal(t, users[0].Password, Password1)

	assert.Equal(t, users[1].ID, UserId2)
	assert.Equal(t, users[1].Password, Password2)

	assert.Equal(t, users[2].ID, UserId3)
	assert.Equal(t, users[2].Password, Password3)
}

func TestGetUserById(t *testing.T) {
	user, err := userRepositoryComponentTest.repository.GetById(context.TODO(), UserId1)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, user, nil)
	assert.Equal(t, user.ID, UserId1)
	assert.Equal(t, user.Password, Password1)
}
