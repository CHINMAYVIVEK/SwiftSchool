package admissions

import (
	"context"
	"swiftschool/domain"

	"github.com/google/uuid"
)

// ========================= CREATE ENQUIRY =========================

// SERVICE
func (s *Service) CreateEnquiry(ctx context.Context, arg domain.AdmissionEnquiry) (*domain.AdmissionEnquiry, error) {
	enquiry, err := s.repo.CreateEnquiry(ctx, arg)
	if err != nil {
		return nil, err
	}
	return enquiry, nil
}

// REPOSITORY
func (r *Repository) CreateEnquiry(ctx context.Context, arg domain.AdmissionEnquiry) (*domain.AdmissionEnquiry, error) {
	return nil, nil
}

// ========================= LIST ENQUIRIES =========================

// SERVICE
func (s *Service) ListEnquiries(ctx context.Context, instituteID uuid.UUID) ([]*domain.AdmissionEnquiry, error) {
	enquiries, err := s.repo.ListEnquiries(ctx, instituteID)
	if err != nil {
		return nil, err
	}
	return enquiries, nil
}

// REPOSITORY
func (r *Repository) ListEnquiries(ctx context.Context, instituteID uuid.UUID) ([]*domain.AdmissionEnquiry, error) {
	return nil, nil
}

// ========================= UPDATE ENQUIRY STATUS =========================

// SERVICE
func (s *Service) UpdateEnquiryStatus(ctx context.Context, id, instituteID uuid.UUID, status domain.AdmissionStatus) error {
	return s.repo.UpdateEnquiryStatus(ctx, id, instituteID, status)
}

// REPOSITORY
func (r *Repository) UpdateEnquiryStatus(ctx context.Context, id, instituteID uuid.UUID, status domain.AdmissionStatus) error {
	return nil
}
