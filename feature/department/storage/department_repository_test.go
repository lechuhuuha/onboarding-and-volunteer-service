package storage

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	return gormDB, mock
}

func TestCreateDepartment(t *testing.T) {
	gormDB, mock := setupMockDB(t)

	repo := NewDepartmentRepository(gormDB)

	department := &domain.Department{
		Name:    "HR",
		Address: "123 HR Street",
		Status:  123,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `departments`").WithArgs(department.Name, department.Address, department.Status).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Create(department)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetDepartmentByID(t *testing.T) {
	gormDB, mock := setupMockDB(t)

	repo := NewDepartmentRepository(gormDB)

	department := &domain.Department{
		Name:    "Finance",
		Address: "456 Finance Street",
		Status:  456,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "address", "status"}).
		AddRow(department.Id, department.Name, department.Address, department.Status)

	mock.ExpectQuery("SELECT * FROM `departments` WHERE `departments`.`id` = ?").WithArgs(department.Id).WillReturnRows(rows)

	result, err := repo.GetByID(department.Id)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, department.Name, result.Name)
	assert.Equal(t, department.Address, result.Address)
	assert.Equal(t, department.Status, result.Status)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateDepartment(t *testing.T) {
	gormDB, mock := setupMockDB(t)

	repo := NewDepartmentRepository(gormDB)

	department := &domain.Department{
		Name:    "IT",
		Address: "789 IT Street",
		Status:  789,
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `departments` SET `name`=?,`address`=?,`status`=? WHERE `id` = ?").WithArgs(department.Name, department.Address, department.Status, department.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Update(department)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteDepartment(t *testing.T) {
	gormDB, mock := setupMockDB(t)

	repo := NewDepartmentRepository(gormDB)

	departmentID := uint(1)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `departments` WHERE `departments`.`id` = ?").WithArgs(departmentID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Delete(departmentID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
