package core

import (
	"context"
	"swiftschool/domain"
	"swiftschool/internal/database"

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
	db *database.Database
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{db: db}
}

//////////////////////////////////////////////////////
//                     SERVICE                      //
//////////////////////////////////////////////////////

type Service struct {
	repo RepositoryInterface
}

func NewService(db *database.Database) *Service {
	return &Service{
		repo: NewRepository(db),
	}
}

//////////////////////////////////////////////////////
//                REPOSITORY INTERFACE              //
//////////////////////////////////////////////////////

type RepositoryInterface interface {
	// ========================= INSTITUTE =========================
	CreateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error)
	GetInstituteById(ctx context.Context, id uuid.UUID) (*domain.Institute, error)
	GetInstituteByCode(ctx context.Context, code string) (*domain.Institute, error)
	UpdateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error)
	DeleteInstitute(ctx context.Context, id uuid.UUID) error
	ListInstitutes(ctx context.Context) ([]*domain.Institute, error)

	// ========================= CLASS =========================
	CreateClass(ctx context.Context, arg domain.Class) (*domain.Class, error)
	UpdateClass(ctx context.Context, arg domain.Class) (*domain.Class, error)
	DeleteClass(ctx context.Context, id uuid.UUID) error
	ListClasses(ctx context.Context, instituteID uuid.UUID) ([]*domain.Class, error)

	// ========================= ACADEMIC SESSION =========================
	CreateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error)
	UpdateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error)
	GetActiveSession(ctx context.Context, instituteID uuid.UUID) (*domain.AcademicSession, error)
	ListAcademicSessions(ctx context.Context, instituteID uuid.UUID) ([]*domain.AcademicSession, error)

	// ========================= DEPARTMENT =========================
	CreateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error)
	UpdateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error)
	DeleteDepartment(ctx context.Context, instituteID, id uuid.UUID) error
	ListDepartments(ctx context.Context, instituteID uuid.UUID) ([]*domain.Department, error)

	// ========================= EMPLOYEE =========================
	CreateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error)
	UpdateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error)
	DeleteEmployee(ctx context.Context, instituteID, id uuid.UUID) error
	GetEmployeeById(ctx context.Context, instituteID, id uuid.UUID) (*domain.Employee, error)
	GetEmployeeFullProfile(ctx context.Context, instituteID, id uuid.UUID) (*domain.Employee, error)
	ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]*domain.Employee, error)

	// ========================= STUDENT =========================
	CreateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error)
	UpdateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error)
	DeleteStudent(ctx context.Context, instituteID, id uuid.UUID) error
	GetStudentFullProfile(ctx context.Context, instituteID, id uuid.UUID) (*domain.Student, error)
	SearchStudents(ctx context.Context, instituteID uuid.UUID, query string) ([]*domain.Student, error)
	ListStudentsByClass(ctx context.Context, instituteID, classID uuid.UUID) ([]*domain.Student, error)

	// ========================= GUARDIAN =========================	// Guardians
	CreateGuardian(ctx context.Context, arg domain.Guardian) (*domain.Guardian, error)
	LinkStudentGuardian(ctx context.Context, studentID, guardianID uuid.UUID, relationship string, isPrimary bool) error

	// ========================= ADDRESS =========================
	CreateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error)
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	// ========================= INSTITUTE =========================
	CreateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error)
	GetInstituteById(ctx context.Context, id uuid.UUID) (*domain.Institute, error)
	GetInstituteByCode(ctx context.Context, code string) (*domain.Institute, error)
	UpdateInstitute(ctx context.Context, arg domain.Institute) (*domain.Institute, error)
	DeleteInstitute(ctx context.Context, id uuid.UUID) error
	ListInstitutes(ctx context.Context) ([]*domain.Institute, error)

	// ========================= CLASS =========================
	CreateClass(ctx context.Context, arg domain.Class) (*domain.Class, error)
	UpdateClass(ctx context.Context, arg domain.Class) (*domain.Class, error)
	DeleteClass(ctx context.Context, id uuid.UUID) error
	ListClasses(ctx context.Context, instituteID uuid.UUID) ([]*domain.Class, error)
	ListStudentsByClass(ctx context.Context, instituteID, classID uuid.UUID) ([]*domain.Student, error)

	// ========================= ACADEMIC SESSION =========================
	CreateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error)
	UpdateAcademicSession(ctx context.Context, arg domain.AcademicSession) (*domain.AcademicSession, error)
	GetActiveSession(ctx context.Context, instituteID uuid.UUID) (*domain.AcademicSession, error)
	ListAcademicSessions(ctx context.Context, instituteID uuid.UUID) ([]*domain.AcademicSession, error)

	// ========================= DEPARTMENT =========================
	CreateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error)
	UpdateDepartment(ctx context.Context, arg domain.Department) (*domain.Department, error)
	DeleteDepartment(ctx context.Context, instituteID, id uuid.UUID) error
	ListDepartments(ctx context.Context, instituteID uuid.UUID) ([]*domain.Department, error)

	// ========================= EMPLOYEE =========================
	CreateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error)
	UpdateEmployee(ctx context.Context, arg domain.Employee) (*domain.Employee, error)
	DeleteEmployee(ctx context.Context, instituteID, id uuid.UUID) error
	GetEmployeeById(ctx context.Context, instituteID, id uuid.UUID) (*domain.Employee, error)
	GetEmployeeFullProfile(ctx context.Context, instituteID, id uuid.UUID) (*domain.Employee, error)
	ListEmployees(ctx context.Context, instituteID uuid.UUID) ([]*domain.Employee, error)

	// ========================= STUDENT =========================
	CreateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error)
	UpdateStudent(ctx context.Context, arg domain.Student) (*domain.Student, error)
	DeleteStudent(ctx context.Context, instituteID, id uuid.UUID) error
	GetStudentFullProfile(ctx context.Context, instituteID, id uuid.UUID) (*domain.Student, error)
	SearchStudents(ctx context.Context, instituteID uuid.UUID, query string) ([]*domain.Student, error)

	// ========================= GUARDIAN =========================
	CreateGuardian(ctx context.Context, arg domain.Guardian) (*domain.Guardian, error)
	LinkStudentGuardian(ctx context.Context, studentID, guardianID uuid.UUID, relationship string, isPrimary bool) error

	// ========================= ADDRESS =========================
	CreateAddress(ctx context.Context, arg domain.Address) (*domain.Address, error)
}
