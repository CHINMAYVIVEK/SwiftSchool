package core

import (
	"context"
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"

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
	CreateClass(ctx context.Context, arg db.CreateClassParams) (db.CoreClass, error)
	CreateAcademicSession(ctx context.Context, arg db.CreateAcademicSessionParams) (db.CoreAcademicSession, error)
	CreateAddress(ctx context.Context, arg db.CreateAddressParams) (db.CoreAddress, error)
	CreateDepartment(ctx context.Context, arg db.CreateDepartmentParams) (db.CoreDepartment, error)
	CreateEmployee(ctx context.Context, arg db.CreateEmployeeParams) (db.CoreEmployee, error)
	CreateGuardian(ctx context.Context, arg db.CreateGuardianParams) (db.CoreGuardian, error)
	CreateStudent(ctx context.Context, arg db.CreateStudentParams) (db.CoreStudent, error)
	DeleteClass(ctx context.Context, arg db.DeleteClassParams) error
	DeleteDepartment(ctx context.Context, arg db.DeleteDepartmentParams) error
	DeleteEmployee(ctx context.Context, arg db.DeleteEmployeeParams) error
	DeleteInstitute(ctx context.Context, arg db.DeleteInstituteParams) error
	DeleteStudent(ctx context.Context, arg db.DeleteStudentParams) error
	GetActiveSession(ctx context.Context, instituteID uuid.UUID) (db.CoreAcademicSession, error)
	GetEmployeeById(ctx context.Context, arg db.GetEmployeeByIdParams) (db.CoreEmployee, error)
	GetEmployeeFullProfile(ctx context.Context, arg db.GetEmployeeFullProfileParams) (db.GetEmployeeFullProfileRow, error)
	GetInstituteByCode(ctx context.Context, code string) (*domain.Institute, error)
	GetInstituteById(ctx context.Context, id uuid.UUID) (*domain.Institute, error)
	GetStudentFullProfile(ctx context.Context, arg db.GetStudentFullProfileParams) (db.GetStudentFullProfileRow, error)
	LinkStudentGuardian(ctx context.Context, arg db.LinkStudentGuardianParams) error
	ListAcademicSessions(ctx context.Context, instituteID uuid.UUID) ([]db.CoreAcademicSession, error)
	ListClasses(ctx context.Context, arg db.ListClassesParams) ([]db.ListClassesRow, error)
	ListDepartments(ctx context.Context, instituteID uuid.NullUUID) ([]db.CoreDepartment, error)
	ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]db.ListEmployeesRow, error)
	ListInstitutes(ctx context.Context) ([]domain.Institute, error)
	ListStudentsByClass(ctx context.Context, arg db.ListStudentsByClassParams) ([]db.CoreStudent, error)
	SearchStudents(ctx context.Context, arg db.SearchStudentsParams) ([]db.CoreStudent, error)
	UpdateAcademicSession(ctx context.Context, arg db.UpdateAcademicSessionParams) (db.CoreAcademicSession, error)
	UpdateClass(ctx context.Context, arg db.UpdateClassParams) (db.CoreClass, error)
	UpdateDepartment(ctx context.Context, arg db.UpdateDepartmentParams) (db.CoreDepartment, error)
	UpdateEmployee(ctx context.Context, arg db.UpdateEmployeeParams) (db.CoreEmployee, error)
	UpdateInstitute(ctx context.Context, arg db.UpdateInstituteParams) (*domain.Institute, error)
	UpdateStudent(ctx context.Context, arg db.UpdateStudentParams) (db.CoreStudent, error)
}

// ------------------ SERVICE INTERFACE ------------------

type ServiceInterface interface {
	CreateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error)
	CreateClass(ctx context.Context, arg domain.Class) (domain.Class, error)
	CreateAcademicSession(ctx context.Context, arg domain.AcademicSession) (domain.AcademicSession, error)
	CreateAddress(ctx context.Context, arg domain.Address) (domain.Address, error)
	CreateDepartment(ctx context.Context, arg domain.Department) (domain.Department, error)
	CreateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error)
	CreateGuardian(ctx context.Context, arg domain.Guardian) (domain.Guardian, error)
	CreateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error)
	DeleteClass(ctx context.Context, id uuid.UUID) error
	DeleteDepartment(ctx context.Context, id uuid.UUID) error
	DeleteEmployee(ctx context.Context, id uuid.UUID) error
	DeleteInstitute(ctx context.Context, id uuid.UUID) error
	DeleteStudent(ctx context.Context, id uuid.UUID) error
	GetActiveSession(ctx context.Context, instituteID uuid.UUID) (domain.AcademicSession, error)
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
	UpdateAcademicSession(ctx context.Context, arg domain.AcademicSession) (domain.AcademicSession, error)
	UpdateClass(ctx context.Context, arg domain.Class) (domain.Class, error)
	UpdateDepartment(ctx context.Context, arg domain.Department) (domain.Department, error)
	UpdateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error)
	UpdateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error)
	UpdateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error)
}
