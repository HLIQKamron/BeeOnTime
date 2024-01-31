package handlers

import (
	"github.com/BeeOntime/config"
	"github.com/BeeOntime/storage"
)

type handlerV1 struct {
	cfg     config.Config
	storage storage.StorageI
}
type Handler struct {
	Cfg      config.Config
	Storage storage.StorageI
}

// New ...
func NewHandler(c *Handler) *handlerV1 {
	return &handlerV1{
		cfg:     c.Cfg,
		storage: c.Storage,
	}
}
