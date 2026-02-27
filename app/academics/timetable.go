package academics

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// ========================= SERVICE =========================

func (s *Service) CreateTimetableEntry(ctx context.Context, arg domain.TimetableEntry) (*domain.TimetableEntry, error) {
	timeTableEntry, err := s.repo.CreateTimetableEntry(ctx, arg)
	if err != nil {
		return nil, err
	}
	return timeTableEntry, nil
}

func (s *Service) GetClassTimetable(ctx context.Context, instituteID, classID uuid.UUID, day domain.DayOfWeek) ([]*domain.TimetableEntry, error) {
	timeTableEntries, err := s.repo.GetClassTimetable(ctx, instituteID, classID, day)
	if err != nil {
		return nil, err
	}
	return timeTableEntries, nil
}

// ========================= REPOSITORY =========================

func (r *Repository) CreateTimetableEntry(ctx context.Context, arg domain.TimetableEntry) (*domain.TimetableEntry, error) {
	return nil, nil
}

func (r *Repository) GetClassTimetable(ctx context.Context, instituteID, classID uuid.UUID, day domain.DayOfWeek) ([]*domain.TimetableEntry, error) {
	return nil, nil
}
