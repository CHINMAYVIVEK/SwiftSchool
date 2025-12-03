package finance

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
	// ========================= FEE MANAGEMENT =========================
	CreateFeeHead(ctx context.Context, arg domain.FeeHead) (*domain.FeeHead, error)
	ListFeeHeads(ctx context.Context, instituteID uuid.UUID) ([]*domain.FeeHead, error)

	CreateFeeStructure(ctx context.Context, arg domain.FeeStructure) (*domain.FeeStructure, error)
	ListFeeStructures(ctx context.Context, instituteID, sessionID uuid.UUID) ([]*domain.FeeStructure, error)

	CreateFineRule(ctx context.Context, arg domain.FineRule) (*domain.FineRule, error)

	// ========================= INVOICES =========================
	CreateInvoice(ctx context.Context, arg domain.Invoice) (*domain.Invoice, error)
	CreateInvoiceItem(ctx context.Context, arg domain.InvoiceItem) (*domain.InvoiceItem, error)
	GetInvoiceById(ctx context.Context, id, instituteID uuid.UUID) (*domain.Invoice, error)
	GetInvoiceWithItems(ctx context.Context, id, instituteID uuid.UUID) (*domain.Invoice, []*domain.InvoiceItem, error)
	ListStudentInvoices(ctx context.Context, instituteID, studentID uuid.UUID) ([]*domain.Invoice, error)
	UpdateInvoiceStatus(ctx context.Context, id, instituteID uuid.UUID, amount float64, status domain.SaaSInvoiceStatus) error
	GetOverdueInvoices(ctx context.Context, instituteID uuid.UUID) ([]*domain.Invoice, error)

	// ========================= TRANSACTIONS =========================
	CreateTransaction(ctx context.Context, arg domain.Transaction) (*domain.Transaction, error)

	// ========================= ACCOUNTING (GL) =========================
	CreateAccount(ctx context.Context, arg domain.Account) (*domain.Account, error)
	ListAccounts(ctx context.Context, instituteID uuid.UUID) ([]*domain.Account, error)
	CreateJournalEntry(ctx context.Context, arg domain.JournalEntry) (*domain.JournalEntry, error)
	CreateJournalItem(ctx context.Context, arg domain.JournalItem) (*domain.JournalItem, error)
	GetAccountBalance(ctx context.Context, instituteID, accountID uuid.UUID) (float64, error)

	// ========================= PROCUREMENT =========================
	CreateVendor(ctx context.Context, arg domain.Vendor) (*domain.Vendor, error)
	ListVendors(ctx context.Context, instituteID uuid.UUID) ([]*domain.Vendor, error)
	CreatePurchaseOrder(ctx context.Context, arg domain.PurchaseOrder) (*domain.PurchaseOrder, error)
	AddPurchaseItem(ctx context.Context, arg domain.PurchaseItem) (*domain.PurchaseItem, error)
	UpdatePurchaseStatus(ctx context.Context, id, instituteID uuid.UUID, status domain.PurchaseStatus) error
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	// ========================= FEE MANAGEMENT =========================
	CreateFeeHead(ctx context.Context, arg domain.FeeHead) (*domain.FeeHead, error)
	ListFeeHeads(ctx context.Context, instituteID uuid.UUID) ([]*domain.FeeHead, error)

	CreateFeeStructure(ctx context.Context, arg domain.FeeStructure) (*domain.FeeStructure, error)
	ListFeeStructures(ctx context.Context, instituteID, sessionID uuid.UUID) ([]*domain.FeeStructure, error)

	CreateFineRule(ctx context.Context, arg domain.FineRule) (*domain.FineRule, error)

	// ========================= INVOICES =========================
	CreateInvoice(ctx context.Context, arg domain.Invoice) (*domain.Invoice, error)
	CreateInvoiceItem(ctx context.Context, arg domain.InvoiceItem) (*domain.InvoiceItem, error)
	GetInvoiceById(ctx context.Context, id, instituteID uuid.UUID) (*domain.Invoice, error)
	GetInvoiceWithItems(ctx context.Context, id, instituteID uuid.UUID) (*domain.Invoice, []*domain.InvoiceItem, error)
	ListStudentInvoices(ctx context.Context, instituteID, studentID uuid.UUID) ([]*domain.Invoice, error)
	UpdateInvoiceStatus(ctx context.Context, id, instituteID uuid.UUID, amount float64, status domain.SaaSInvoiceStatus) error
	GetOverdueInvoices(ctx context.Context, instituteID uuid.UUID) ([]*domain.Invoice, error)

	// ========================= TRANSACTIONS =========================
	CreateTransaction(ctx context.Context, arg domain.Transaction) (*domain.Transaction, error)

	// ========================= ACCOUNTING (GL) =========================
	CreateAccount(ctx context.Context, arg domain.Account) (*domain.Account, error)
	ListAccounts(ctx context.Context, instituteID uuid.UUID) ([]*domain.Account, error)
	CreateJournalEntry(ctx context.Context, arg domain.JournalEntry) (*domain.JournalEntry, error)
	CreateJournalItem(ctx context.Context, arg domain.JournalItem) (*domain.JournalItem, error)
	GetAccountBalance(ctx context.Context, instituteID, accountID uuid.UUID) (float64, error)

	// ========================= PROCUREMENT =========================
	CreateVendor(ctx context.Context, arg domain.Vendor) (*domain.Vendor, error)
	ListVendors(ctx context.Context, instituteID uuid.UUID) ([]*domain.Vendor, error)
	CreatePurchaseOrder(ctx context.Context, arg domain.PurchaseOrder) (*domain.PurchaseOrder, error)
	AddPurchaseItem(ctx context.Context, arg domain.PurchaseItem) (*domain.PurchaseItem, error)
	UpdatePurchaseStatus(ctx context.Context, id, instituteID uuid.UUID, status domain.PurchaseStatus) error
}
