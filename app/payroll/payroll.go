package payroll

import (
	"context"
	"swiftschool/domain"
	"swiftschool/helper"
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
	CreateSalaryComponent(ctx context.Context, arg domain.SalaryComponent) (*domain.SalaryComponent, error)
	AssignSalaryConfig(ctx context.Context, arg domain.EmployeeSalaryConfig) (*domain.EmployeeSalaryConfig, error)
	CreatePayslip(ctx context.Context, arg domain.Payslip) (*domain.Payslip, error)
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	CreateSalaryComponent(ctx context.Context, arg domain.SalaryComponent) (*domain.SalaryComponent, error)
	AssignSalaryConfig(ctx context.Context, arg domain.EmployeeSalaryConfig) (*domain.EmployeeSalaryConfig, error)
	CreatePayslip(ctx context.Context, arg domain.Payslip) (*domain.Payslip, error)
}
