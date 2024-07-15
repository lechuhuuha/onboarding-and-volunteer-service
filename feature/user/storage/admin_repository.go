package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}
func (r *AdminRepository) GetListPendingRequest() ([]*domain.Request, string) {
	var listRequest []*domain.Request
	result := r.DB.Where("status = ?", 0).Find(&listRequest)
	if result.Error != nil {
		return nil, result.Error.Error()
	}
	if len(listRequest) == 0 {
		return nil, "No request found"
	}
	return listRequest, ""
}

func (r *AdminRepository) GetRequestByID(id int) (*domain.Request, string) {
	var request domain.Request
	result := r.DB.Where("id = ? and status = 0", id).First(&request)
	if result.Error != nil {
		return nil, result.Error.Error()
	}
	return &request, ""
}

func (r *AdminRepository) ApproveRequest(id int, verifier_id int) string {
	result := r.DB.Model(&domain.Request{}).Where("id = ?", id).Update("status", 1).Update("verifier_id", verifier_id)
	if result.Error != nil {
		return result.Error.Error()
	}
	return "Approve request success"
}
func (r *AdminRepository) RejectRequest(id int, verifier_id int) string {
	result := r.DB.Model(&domain.Request{}).Where("id = ?", id).Update("status", 2).Update("verifier_id", verifier_id)
	if result.Error != nil {
		return result.Error.Error()
	}
	return "Reject request success"
}
func (r *AdminRepository) AddRejectNotes(id int, notes string) string {
	result := r.DB.Model(&domain.Request{}).Where("id = ?", id).Update("reject_notes", notes)
	if result.Error != nil {
		return result.Error.Error()
	}
	return "Add reject notes success"
}
func (r *AdminRepository) DeleteRequest(id int) string {
	result := r.DB.Where("id = ?", id).Delete(&domain.Request{})
	if result.Error != nil {
		return result.Error.Error()
	}
	return "Delete request success"
}
