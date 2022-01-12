package test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"regexp"

	"gorm.io/gorm"
	"testing"
)

type UserRepositoryTest struct {
	mock       sqlmock.Sqlmock
	mockDB     *gorm.DB
	repository repositories.UserRepository
}

func NewUserRepositoryTest(db *gorm.DB, mock sqlmock.Sqlmock,
	userRepository repositories.UserRepository) *UserRepositoryTest {
	return &UserRepositoryTest{
		mockDB:     db,
		mock:       mock,
		repository: userRepository,
	}
}

const (
	IdField       string = "id"
	PasswordField string = "password"
)

var (
	UserId   = "USER-ID-HERE"
	Password = "USER-PASSWORD-HERE"
)

var (
	UserId1   = "USER-ID-1"
	Password1 = "PASSWORD-1"
	UserId2   = "USER-ID-2"
	Password2 = "PASSWORD-2"
	UserId3   = "USER-ID-3"
	Password3 = "PASSWORD-3"
)

var userRepositoryTest *UserRepositoryTest

var userField = []string{IdField, PasswordField}

func init() {
	logrus.Infof("Start Initializing User Repository Test Successfully")
	db, mock, err := sqlmock.New()
	if err != nil {
		logrus.Errorf("Init SQL Mock error: %s", err.Error())
		return
	}

	dialect := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	mockDB, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		logrus.Errorf("Open Gorm Mock error: %s", err.Error())
		return
	}

	userRepository := repositories.NewUserRepository(mockDB)
	userRepositoryTest = NewUserRepositoryTest(mockDB, mock, userRepository)
	logrus.Infof("End Initializing User Repository Test Successfully")
}

func TestGetAllUser(t *testing.T) {
	var query = `SELECT * FROM "users"`
	var rows = sqlmock.NewRows(userField).AddRow(UserId1, Password1).
		AddRow(UserId2, Password2).
		AddRow(UserId3, Password3)

	userRepositoryTest.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
	configurations, err := userRepositoryTest.repository.GetAll(context.TODO())

	assert.Equal(t, err, nil)
	assert.Equal(t, len(configurations), 3)

	if err := userRepositoryTest.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}
}

func TestGetUserById(t *testing.T) {
	var query = `SELECT * FROM "users" WHERE id = $1`

	var row = sqlmock.NewRows(userField).AddRow(UserId, Password)
	userRepositoryTest.mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(UserId).
		WillReturnRows(row)

	user, err := userRepositoryTest.repository.GetById(context.TODO(), UserId)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, user, nil)
	assert.Equal(t, user.ID, UserId)
	assert.Equal(t, user.Password, Password)

	if err := userRepositoryTest.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}
}
