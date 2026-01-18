package library

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
	CreateBookCategory(ctx context.Context, arg domain.BookCategory) (*domain.BookCategory, error)
	CreateBook(ctx context.Context, arg domain.Book) (*domain.Book, error)
	ListBooks(ctx context.Context, instituteID uuid.UUID) ([]*domain.Book, error)
	SearchBooks(ctx context.Context, instituteID uuid.UUID, query string) ([]*domain.Book, error)

	IssueBook(ctx context.Context, arg domain.BookIssue) (*domain.BookIssue, error)
	ReturnBook(ctx context.Context, id, instituteID uuid.UUID, fineAmount float64) error
}

//////////////////////////////////////////////////////
//                 SERVICE INTERFACE                //
//////////////////////////////////////////////////////

type ServiceInterface interface {
	CreateBookCategory(ctx context.Context, arg domain.BookCategory) (*domain.BookCategory, error)
	CreateBook(ctx context.Context, arg domain.Book) (*domain.Book, error)
	ListBooks(ctx context.Context, instituteID uuid.UUID) ([]*domain.Book, error)
	SearchBooks(ctx context.Context, instituteID uuid.UUID, query string) ([]*domain.Book, error)

	IssueBook(ctx context.Context, arg domain.BookIssue) (*domain.BookIssue, error)
	ReturnBook(ctx context.Context, id, instituteID uuid.UUID, fineAmount float64) error
}
