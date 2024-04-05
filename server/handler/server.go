package handler

import (
	"CarCatalog/model"
	"CarCatalog/storage"
)

type Server struct {
	Config  model.Config
	Storage *storage.Storage
}
