package core

type CoreHandler struct {
	service *CoreService
}

func NewHandler(service *CoreService) *CoreHandler {
	return &CoreHandler{service: service}
}
