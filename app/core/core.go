package core

import (
	"context"
	"swiftschool/domain"
	"swiftschool/helper"

	"github.com/google/uuid"
)

// ------------------ HANDLER ------------------

type Handler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{service: service}
}

// ------------------ REPOSITORY ------------------

type Repository struct {
	db *helper.PostgresWrapper
}

func NewRepository(db *helper.PostgresWrapper) *Repository {
	return &Repository{db: db}
}

// ------------------ SERVICE ------------------

type Service struct {
	repo RepositoryInterface
}

func NewService(db *helper.PostgresWrapper) *Service {
	return &Service{
		repo: NewRepository(db),
	}
}

// ------------------ REPOSITORY INTERFACE ------------------

type RepositoryInterface interface {
	CreateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error)
	CreateClass(ctx context.Context, arg domain.Class) (*domain.Class, error)
	CreateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error)
	CreateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error)
	CreateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error)
	CreateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error)
	CreateGuardian(ctx context.Context, arg domain.Guardian) (*domain.Guardian, error)
	CreateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error)
	DeleteClass(ctx context.Context, id uuid.UUID) error
	DeleteDepartment(ctx context.Context, id uuid.UUID) error
	DeleteEmployee(ctx context.Context, id uuid.UUID) error
	DeleteInstitute(ctx context.Context, id uuid.UUID) error
	DeleteStudent(ctx context.Context, id uuid.UUID) error
	GetActiveSession(ctx context.Context, instituteID uuid.UUID) (*domain.AcademicSession, error)
	GetEmployeeById(ctx context.Context, id uuid.UUID) (*domain.Employee, error)
	GetEmployeeFullProfile(ctx context.Context, id uuid.UUID) (*domain.Employee, error)
	GetInstituteByCode(ctx context.Context, code string) (*domain.Institute, error)
	GetInstituteById(ctx context.Context, id uuid.UUID) (*domain.Institute, error)
	GetStudentFullProfile(ctx context.Context, id uuid.UUID) (*domain.Student, error)
	LinkStudentGuardian(ctx context.Context, studentID, guardianID uuid.UUID) error
	ListAcademicSessions(ctx context.Context, instituteID uuid.UUID) ([]domain.AcademicSession, error)
	ListClasses(ctx context.Context, instituteID uuid.UUID) ([]domain.Class, error)
	ListDepartments(ctx context.Context, instituteID uuid.UUID) ([]domain.Department, error)
	ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]domain.Employee, error)
	ListInstitutes(ctx context.Context) ([]domain.Institute, error)
	ListStudentsByClass(ctx context.Context, classID uuid.UUID) ([]domain.Student, error)
	SearchStudents(ctx context.Context, query string) ([]domain.Student, error)
	UpdateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error)
	UpdateClass(ctx context.Context, arg domain.Class) (*domain.Class, error)
	UpdateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error)
	UpdateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error)
	UpdateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error)
	UpdateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error)
}

// ------------------ SERVICE INTERFACE ------------------

type ServiceInterface interface {
	CreateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error)
	CreateClass(ctx context.Context, arg domain.Class) (*domain.Class, error)
	CreateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error)
	CreateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error)
	CreateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error)
	CreateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error)
	CreateGuardian(ctx context.Context, arg domain.Guardian) (*domain.Guardian, error)
	CreateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error)
	DeleteClass(ctx context.Context, id uuid.UUID) error
	DeleteDepartment(ctx context.Context, id uuid.UUID) error
	DeleteEmployee(ctx context.Context, id uuid.UUID) error
	DeleteInstitute(ctx context.Context, id uuid.UUID) error
	DeleteStudent(ctx context.Context, id uuid.UUID) error
	GetActiveSession(ctx context.Context, instituteID uuid.UUID) (*domain.AcademicSession, error)
	GetEmployeeById(ctx context.Context, id uuid.UUID) (*domain.Employee, error)
	GetEmployeeFullProfile(ctx context.Context, id uuid.UUID) (*domain.Employee, error)
	GetInstituteByCode(ctx context.Context, code string) (*domain.Institute, error)
	GetInstituteById(ctx context.Context, id uuid.UUID) (*domain.Institute, error)
	GetStudentFullProfile(ctx context.Context, id uuid.UUID) (*domain.Student, error)
	LinkStudentGuardian(ctx context.Context, studentID, guardianID uuid.UUID) error
	ListAcademicSessions(ctx context.Context, instituteID uuid.UUID) ([]domain.AcademicSession, error)
	ListClasses(ctx context.Context, instituteID uuid.UUID) ([]domain.Class, error)
	ListDepartments(ctx context.Context, instituteID uuid.UUID) ([]domain.Department, error)
	ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]domain.Employee, error)
	ListInstitutes(ctx context.Context) ([]domain.Institute, error)
	ListStudentsByClass(ctx context.Context, classID uuid.UUID) ([]domain.Student, error)
	SearchStudents(ctx context.Context, query string) ([]domain.Student, error)
	UpdateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error)
	UpdateClass(ctx context.Context, arg domain.Class) (*domain.Class, error)
	UpdateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error)
	UpdateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error)
	UpdateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error)
	UpdateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error)
}
