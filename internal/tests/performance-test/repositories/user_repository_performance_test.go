package test

import (
	"context"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
	"testing"
)

type UserRepositoryPerformanceTest struct {
	repository repositories.UserRepository
}

func NewUserRepositoryPerformanceTest(userRepository repositories.UserRepository) *UserRepositoryPerformanceTest {
	return &UserRepositoryPerformanceTest{
		repository: userRepository,
	}
}

var (
	UserId1   = "firstUser"
)

var userRepositoryPerformanceTest *UserRepositoryPerformanceTest

func init() {
	logrus.Infof("Start Initializing User Repository Test Successfully")
	db, err := gorm.Open(sqlite.Open("../../../../data.db"), &gorm.Config{})
	if err != nil {
		logrus.Errorf("Open Gorm Mock error: %s", err.Error())
		return
	}

	userRepository := repositories.NewUserRepository(db)
	userRepositoryPerformanceTest = NewUserRepositoryPerformanceTest(userRepository)
	logrus.Infof("End Initializing User Repository Test Successfully")
}

func BenchmarkGetAllUser(b *testing.B) {
	for x := 0; x < b.N; x++ {
		_, _ = userRepositoryPerformanceTest.repository.GetAll(context.TODO())
	}
}

func BenchmarkGetUserById(b *testing.B) {
	for x := 0; x < b.N; x++ {
		_, _ = userRepositoryPerformanceTest.repository.GetById(context.TODO(), UserId1)
	}
}
