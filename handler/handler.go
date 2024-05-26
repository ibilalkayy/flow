package handler

import "github.com/ibilalkayy/flow/interfaces"

type Handler struct {
	Deps interfaces.Dependencies
}

func NewHandler(deps interfaces.Dependencies) *Handler {
	return &Handler{Deps: deps}
}
