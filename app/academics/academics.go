package academics

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
//               REPOSITORY INTERFACE               //
//////////////////////////////////////////////////////

type RepositoryInterface interface {
	// ========================= SUBJECTS =========================
	CreateSubject(ctx context.Context, arg domain.Subject) (*domain.Subject, error)
	ListSubjects(ctx context.Context, instituteID uuid.UUID) ([]*domain.Subject, error)

	// ========================= CLASS PERIODS =========================
	CreateClassPeriod(ctx context.Context, arg domain.ClassPeriod) (*domain.ClassPeriod, error)
	ListClassPeriods(ctx context.Context, instituteID uuid.UUID) ([]*domain.ClassPeriod, error)

	// ========================= TIMETABLE =========================
	CreateTimetableEntry(ctx context.Context, arg domain.TimetableEntry) (*domain.TimetableEntry, error)
	GetClassTimetable(ctx context.Context, instituteID, classID uuid.UUID, day domain.DayOfWeek) ([]*domain.TimetableEntry, error)
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	// ========================= SUBJECTS =========================
	CreateSubject(ctx context.Context, arg domain.Subject) (*domain.Subject, error)
	ListSubjects(ctx context.Context, instituteID uuid.UUID) ([]*domain.Subject, error)

	// ========================= CLASS PERIODS =========================
	CreateClassPeriod(ctx context.Context, arg domain.ClassPeriod) (*domain.ClassPeriod, error)
	ListClassPeriods(ctx context.Context, instituteID uuid.UUID) ([]*domain.ClassPeriod, error)

	// ========================= TIMETABLE =========================
	CreateTimetableEntry(ctx context.Context, arg domain.TimetableEntry) (*domain.TimetableEntry, error)
	GetClassTimetable(ctx context.Context, instituteID, classID uuid.UUID, day domain.DayOfWeek) ([]*domain.TimetableEntry, error)
}
