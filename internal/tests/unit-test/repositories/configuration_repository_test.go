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

type ConfigurationRepositoryTest struct {
	mock       sqlmock.Sqlmock
	mockDB     *gorm.DB
	repository repositories.ConfigurationRepository
}

func NewConfigurationRepositoryTest(db *gorm.DB, mock sqlmock.Sqlmock,
	configurationRepository repositories.ConfigurationRepository) *ConfigurationRepositoryTest {
	return &ConfigurationRepositoryTest{
		mockDB:     db,
		mock:       mock,
		repository: configurationRepository,
	}
}

var configurationRepositoryTest *ConfigurationRepositoryTest

var (
	ConfigurationId       = "Configuration-Id-Here"
	ConfigurationUserId   = "Configuration-User-Id-Here"
	ConfigurationCapacity = 100
	ConfigurationDate     = "2020-01-01"
)

var (
	ConfigurationId1       = "Configuration-Id-Here"
	ConfigurationUserId1   = "Configuration-User-Id-Here"
	ConfigurationCapacity1 = 100
	ConfigurationDate1     = "2020-01-01"

	ConfigurationId2       = "Configuration-Id-Here-2"
	ConfigurationUserId2   = "Configuration-User-Id-Here-2"
	ConfigurationCapacity2 = 200
	ConfigurationDate2     = "2020-01-02"

	ConfigurationId3       = "Configuration-Id-Here-3"
	ConfigurationUserId3   = "Configuration-User-Id-Here-3"
	ConfigurationCapacity3 = 300
	ConfigurationDate3     = "2020-01-03"
)

var ConfigurationField = []string{"id", "user_id", "capacity", "date"}

func init() {
	logrus.Infof("Start Initializing Configuration Repository Test Successfully")
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

	configurationRepository := repositories.NewConfigurationRepository(mockDB)
	configurationRepositoryTest = NewConfigurationRepositoryTest(mockDB, mock, configurationRepository)
	logrus.Infof("End Initializing Configuration Repository Test Successfully")
}

func TestGetAllConfiguration(t *testing.T) {
	var query = `SELECT * FROM "configurations"`
	var rows = sqlmock.NewRows(ConfigurationField).
		AddRow(ConfigurationId1, ConfigurationUserId1, ConfigurationCapacity1, ConfigurationDate1).
		AddRow(ConfigurationId2, ConfigurationUserId2, ConfigurationCapacity2, ConfigurationDate2).
		AddRow(ConfigurationId3, ConfigurationUserId3, ConfigurationCapacity3, ConfigurationDate3)

	configurationRepositoryTest.mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	configurations, err := configurationRepositoryTest.repository.GetAll(context.TODO())
	assert.Equal(t, err, nil)
	assert.Equal(t, len(configurations), 3)

	if err := configurationRepositoryTest.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("TestGetAllConfiguration ExpectationsWereMet: %s", err)
		t.Fail()
	}
}

func TestGetConfigurationCapacity(t *testing.T) {
	var query = `SELECT * FROM "configurations" WHERE user_id = $1 and date = $2`
	var row = sqlmock.NewRows(ConfigurationField).
		AddRow(ConfigurationId, ConfigurationUserId, ConfigurationCapacity, ConfigurationDate)

	configurationRepositoryTest.mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(ConfigurationId, ConfigurationDate).
		WillReturnRows(row)

	capacity, err := configurationRepositoryTest.repository.GetCapacity(context.TODO(), ConfigurationId, ConfigurationDate)
	assert.Equal(t, err, nil)
	assert.Equal(t, capacity, int64(100))
	if err := configurationRepositoryTest.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("TestGetConfigurationCapacity ExpectationsWereMet: %s", err)
		t.Fail()
	}
}

func TestGetConfigurationById(t *testing.T) {
	var query = `SELECT * FROM "configurations" WHERE id = $1`
	var row = sqlmock.NewRows(ConfigurationField).
		AddRow(ConfigurationId, ConfigurationUserId, ConfigurationCapacity, ConfigurationDate)

	configurationRepositoryTest.mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(ConfigurationId).
		WillReturnRows(row)

	configuration, err := configurationRepositoryTest.repository.GetById(context.TODO(), ConfigurationId)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, configuration,nil)

	if err := configurationRepositoryTest.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("TestGetConfigurationCapacity ExpectationsWereMet: %s", err)
		t.Fail()
	}
}
