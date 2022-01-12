package test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type TaskRepositoryTest struct {
	mock       sqlmock.Sqlmock
	mockDB     *gorm.DB
	repository repositories.TaskRepository
}

func NewTaskRepositoryTest(db *gorm.DB, mock sqlmock.Sqlmock,
	taskRepository repositories.TaskRepository) *TaskRepositoryTest {
	return &TaskRepositoryTest{
		mockDB:     db,
		mock:       mock,
		repository: taskRepository,
	}
}

var (
	TaskId          = "Task-Id-Here"
	TaskUserId      = "Task-User-Id-Here"
	TaskContent     = "Task-Content-Here"
	TaskCreatedDate = "2020-01-01"
)

var (
	TaskId1          = "Task-Id-Here"
	TaskUserId1      = "Task-User-Id-Here"
	TaskContent1     = "Task-Content-Here"
	TaskCreatedDate1 = "2020-01-01"

	TaskId2          = "Task-Id-Here-2"
	TaskUserId2      = "Task-User-Id-Here-2"
	TaskContent2     = "Task-Content-Here-2"
	TaskCreatedDate2 = "2020-01-02"

	TaskId3          = "Task-Id-Here-3"
	TaskUserId3      = "Task-User-Id-Here-3"
	TaskContent3     = "Task-Content-Here-3"
	TaskCreatedDate3 = "2020-01-03"
)

var taskRepositoryTest *TaskRepositoryTest

var TaskField = []string{"id", "content", "user_id", "created_date"}

func init() {
	logrus.Infof("Start Initializing Task Repository Test Successfully")
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

	taskRepository := repositories.NewTaskRepository(mockDB)
	taskRepositoryTest = NewTaskRepositoryTest(mockDB, mock, taskRepository)
	logrus.Infof("End Initializing Task Repository Test Successfully")
}

func TestGetAll(t *testing.T) {
	var query = `SELECT * FROM "tasks"`
	var rows = sqlmock.NewRows(TaskField).
		AddRow(TaskId1, TaskContent1, TaskUserId1, TaskCreatedDate1).
		AddRow(TaskId2, TaskContent2, TaskUserId2, TaskCreatedDate2).
		AddRow(TaskId3, TaskContent3, TaskUserId3, TaskCreatedDate3)

	taskRepositoryTest.mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	tasks, err := taskRepositoryTest.repository.GetAll(context.TODO())
	assert.Equal(t, err, nil)
	assert.Equal(t, len(tasks), 3)

	if err := taskRepositoryTest.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}
}

func TestGetById(t *testing.T) {
	var query = `SELECT * FROM "tasks" WHERE id = $1`
	var row = sqlmock.NewRows(TaskField).
		AddRow(TaskId, TaskUserId, TaskContent,TaskCreatedDate)

	taskRepositoryTest.mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(TaskId).
		WillReturnRows(row)

	task, err := taskRepositoryTest.repository.GetById(context.TODO(), TaskId)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, task, nil)

	if err := taskRepositoryTest.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}
}
