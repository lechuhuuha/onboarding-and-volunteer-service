package dto

type VolunteerCreateDTO struct {
	UserId       uint `json:"user_id" binding:"required"`
	DepartmentId uint `json:"department_id" binding:"required"`
	Status       uint `json:"status" binding:"required"`
}

type VolunteerUpdateDTO struct {
	DepartmentId uint `json:"department_id" binding:"required"`
	Status       uint `json:"status" binding:"required"`
}
