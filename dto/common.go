package dto

// SuccessResponse represents a standard successful API response
type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Operation completed successfully"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents a standard error API response
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"An error occurred"`
	Error   string `json:"error,omitempty" example:"detailed error message"`
}

// PaginationParams represents common pagination query parameters
type PaginationParams struct {
	Page     int `json:"page" example:"1"`
	PageSize int `json:"page_size" example:"10"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Data retrieved successfully"`
	Data       interface{} `json:"data"`
	Page       int         `json:"page" example:"1"`
	PageSize   int         `json:"page_size" example:"10"`
	TotalCount int         `json:"total_count" example:"100"`
	TotalPages int         `json:"total_pages" example:"10"`
}
