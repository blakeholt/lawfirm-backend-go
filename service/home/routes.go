package home

import (
	"lawfirm-go-backend/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/home", h.handleGetHome).Methods("GET")
}

func (h *Handler) handleGetHome(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, "Hello World")
}
