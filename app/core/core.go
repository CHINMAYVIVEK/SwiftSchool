package core

import (
	"context"
	"swiftschool/domain"
	"swiftschool/helper"
)

type Handler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{service: service}
}

type Repository struct {
	db *helper.PostgresWrapper
}

func NewRepository(db *helper.PostgresWrapper) *Repository {
	return &Repository{db: db}
}

type Service struct {
	repo RepositoryInterface
}

func NewService(db *helper.PostgresWrapper) *Service {
	return &Service{
		repo: NewRepository(db),
	}
}

type RepositoryInterface interface {
	CreateInstitute(ctx context.Context, institute domain.Institute) (*domain.BaseUUIDModel, error)
	CreateClass()
	DeleteClass()
	CreateAcademicSession()
	CreateAddress()
	CreateDepartment()
	DeleteDepartment()
	CreateEmployee()
	CreateGuardian()
	CreateStudent()
	LinkStudentGuardian()
}

type ServiceInterface interface {
	CreateInstitute(ctx context.Context, institute domain.Institute) (*domain.BaseUUIDModel, error)
	CreateClass()
	DeleteClass()
	CreateAcademicSession()
	CreateAddress()
	CreateDepartment()
	DeleteDepartment()
	CreateEmployee()
	CreateGuardian()
	CreateStudent()
	LinkStudentGuardian()
}
