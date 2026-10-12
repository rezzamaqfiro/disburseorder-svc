package order

import (
	repo "disburseorder-svc/repo/generated"
)

type Handler struct {
	db *repo.Queries
}

func NewHandler(db *repo.Queries) *Handler {
	return &Handler{db}
}
