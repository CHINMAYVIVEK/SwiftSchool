package subscription

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
	// ========================= PLANS =========================
	CreatePlan(ctx context.Context, arg domain.Plan) (*domain.Plan, error)
	ListPlans(ctx context.Context) ([]*domain.Plan, error)
	UpdatePlan(ctx context.Context, arg domain.Plan) (*domain.Plan, error)

	// ========================= SUBSCRIPTIONS =========================
	CreateSubscription(ctx context.Context, arg domain.Subscription) (*domain.Subscription, error)
	GetActiveSubscription(ctx context.Context, instituteID uuid.UUID) (*domain.Subscription, error)

	// ========================= INVOICES =========================
	CreateSaaSInvoice(ctx context.Context, arg domain.SaaSInvoice) (*domain.SaaSInvoice, error)
	ListSaaSInvoices(ctx context.Context, subscriptionID uuid.UUID) ([]*domain.SaaSInvoice, error)
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	// ========================= PLANS =========================
	CreatePlan(ctx context.Context, arg domain.Plan) (*domain.Plan, error)
	ListPlans(ctx context.Context) ([]*domain.Plan, error)
	UpdatePlan(ctx context.Context, arg domain.Plan) (*domain.Plan, error)

	// ========================= SUBSCRIPTIONS =========================
	CreateSubscription(ctx context.Context, arg domain.Subscription) (*domain.Subscription, error)
	GetActiveSubscription(ctx context.Context, instituteID uuid.UUID) (*domain.Subscription, error)

	// ========================= INVOICES =========================
	CreateSaaSInvoice(ctx context.Context, arg domain.SaaSInvoice) (*domain.SaaSInvoice, error)
	ListSaaSInvoices(ctx context.Context, subscriptionID uuid.UUID) ([]*domain.SaaSInvoice, error)
}
