// department_handler_test.go
package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDepartmentUsecase is a mock implementation of the DepartmentUsecase
type MockDepartmentUsecase struct {
	mock.Mock
}

func (m *MockDepartmentUsecase) CreateDepartment(input dto.DepartmentCreateDTO) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockDepartmentUsecase) GetDepartmentByID(id uint) (*domain.Department, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Department), args.Error(1)
}

func (m *MockDepartmentUsecase) UpdateDepartment(id uint, input dto.DepartmentUpdateDTO) error {
	args := m.Called(id, input)
	return args.Error(0)
}

func (m *MockDepartmentUsecase) DeleteDepartment(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateDepartment(t *testing.T) {
	mockUsecase := new(MockDepartmentUsecase)
	handler := NewDepartmentHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/v1/departments", handler.CreateDepartment)

	input := dto.DepartmentCreateDTO{
		Name:    "HR",
		Address: "123 HR Street",
		Status:  123,
	}
	response := dto.DepartmentCreateDTO{
		Name:    "HR",
		Address: "123 HR Street",
		Status:  123,
	}

	mockUsecase.On("CreateDepartment", input).Return(response, nil)

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/api/v1/departments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var result dto.DepartmentResponseDTO
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, response, result)
	assert.JSONEq(t, `{"Name":"HR","Address":"123 HR Street","Status":123}`, w.Body.String())

	mockUsecase.AssertExpectations(t)
}

func TestGetDepartmentByID(t *testing.T) {
	mockUsecase := new(MockDepartmentUsecase)
	handler := NewDepartmentHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/v1/departments/:id", handler.GetDepartmentByID)

	response := dto.DepartmentResponseDTO{Name: "Test Department"}

	mockUsecase.On("GetDepartmentByID", uint(1)).Return(response, nil)

	req, _ := http.NewRequest("GET", "/api/v1/departments/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var result dto.DepartmentResponseDTO
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, response, result)
	assert.JSONEq(t, `{"ID":1,"Name":"Finance","Address":"456 Finance Street","Status":456}`, w.Body.String())

	mockUsecase.AssertExpectations(t)
}

func TestUpdateDepartment(t *testing.T) {
	mockUsecase := new(MockDepartmentUsecase)
	handler := NewDepartmentHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/api/v1/departments/:id", handler.UpdateDepartment)

	input := dto.DepartmentUpdateDTO{
		Name:    "IT Updated",
		Address: "789 IT Street Updated",
		Status:  789,
	}
	mockUsecase.On("UpdateDepartment", uint(1), input).Return(nil, nil)

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("PUT", "/api/v1/departments/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var result dto.DepartmentResponseDTO
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, nil, result)
	assert.JSONEq(t, `{"ID":1,"Name":"IT Updated","Address":"789 IT Street Updated","Status":"Inactive"}`, w.Body.String())

	mockUsecase.AssertExpectations(t)
}

func TestDeleteDepartment(t *testing.T) {
	mockUsecase := new(MockDepartmentUsecase)
	handler := NewDepartmentHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/api/v1/departments/:id", handler.DeleteDepartment)

	mockUsecase.On("DeleteDepartment", uint(1)).Return(nil)

	req, _ := http.NewRequest("DELETE", "/api/v1/departments/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.Bytes())

	mockUsecase.AssertExpectations(t)
}
