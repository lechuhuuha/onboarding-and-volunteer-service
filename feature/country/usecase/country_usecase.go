package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/storage"
)

// CountryUsecaseInterface defines the methods that any use case implementation must provide.
type CountryUsecaseInterface interface {
	CreateCountry(input dto.CountryCreateDTO) (*domain.Country, error)
	GetCountryByID(id uint) (*domain.Country, error)
	UpdateCountry(id uint, input dto.CountryUpdateDTO) (*domain.Country, error)
	DeleteCountry(id uint) error
}

// CountryUsecase handles the business logic for countries.
type CountryUsecase struct {
	repo storage.CountryRepositoryInterface
}

// NewCountryUsecase creates a new instance of CountryUsecase.
func NewCountryUsecase(repo storage.CountryRepositoryInterface) *CountryUsecase {
	return &CountryUsecase{repo: repo}
}

// CreateCountry creates a new country using the provided DTO.
func (u *CountryUsecase) CreateCountry(input dto.CountryCreateDTO) (*domain.Country, error) {
	country := &domain.Country{
		Name:   input.Name,
		Status: input.Status,
	}
	err := u.repo.Create(country)
	return country, err
}

// GetCountryByID retrieves a country by its ID.
func (u *CountryUsecase) GetCountryByID(id uint) (*domain.Country, error) {
	return u.repo.GetByID(id)
}

// UpdateCountry updates a country using the provided DTO.
func (u *CountryUsecase) UpdateCountry(id uint, input dto.CountryUpdateDTO) (*domain.Country, error) {
	country, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	country.Name = input.Name
	country.Status = input.Status
	err = u.repo.Update(country)
	return country, err
}

// DeleteCountry deletes a country by its ID.
func (u *CountryUsecase) DeleteCountry(id uint) error {
	return u.repo.Delete(id)
}

// Ensure CountryUsecase implements CountryUsecaseInterface
var _ CountryUsecaseInterface = (*CountryUsecase)(nil)
