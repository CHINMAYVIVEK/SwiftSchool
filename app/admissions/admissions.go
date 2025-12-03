package admissions

import (
	"context"
	"swiftschool/domain"
	"swiftschool/helper"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////////
//                     HANDLER                      //
//////////////////////////////////////////////////////

type Handler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{service: service}
}

//////////////////////////////////////////////////////
//                    REPOSITORY                    //
//////////////////////////////////////////////////////

type Repository struct {
	db *helper.PostgresWrapper
}

func NewRepository(db *helper.PostgresWrapper) *Repository {
	return &Repository{db: db}
}

//////////////////////////////////////////////////////
//                     SERVICE                      //
//////////////////////////////////////////////////////

type Service struct {
	repo RepositoryInterface
}

func NewService(db *helper.PostgresWrapper) *Service {
	return &Service{
		repo: NewRepository(db),
	}
}

//////////////////////////////////////////////////////
//               REPOSITORY INTERFACE               //
//////////////////////////////////////////////////////

type RepositoryInterface interface {
	CreateEnquiry(ctx context.Context, arg domain.AdmissionEnquiry) (*domain.AdmissionEnquiry, error)
	ListEnquiries(ctx context.Context, instituteID uuid.UUID) ([]*domain.AdmissionEnquiry, error)
	UpdateEnquiryStatus(ctx context.Context, id, instituteID uuid.UUID, status domain.AdmissionStatus) error
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	CreateEnquiry(ctx context.Context, arg domain.AdmissionEnquiry) (*domain.AdmissionEnquiry, error)
	ListEnquiries(ctx context.Context, instituteID uuid.UUID) ([]*domain.AdmissionEnquiry, error)
	UpdateEnquiryStatus(ctx context.Context, id, instituteID uuid.UUID, status domain.AdmissionStatus) error
}
