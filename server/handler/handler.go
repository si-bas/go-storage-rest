package handler

import "github.com/si-bas/go-storage-rest/service"

type Handler struct {
	clientService service.ClientService
	fileService   service.FileService

	// Public
	ClientService service.ClientService
}

func New(clientService service.ClientService, fileService service.FileService) *Handler {
	return &Handler{
		clientService: clientService,
		fileService:   fileService,

		ClientService: clientService,
	}
}
