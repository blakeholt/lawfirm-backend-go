package client

import (
	"fmt"
	"lawfirm-go-backend/service/auth"
	"lawfirm-go-backend/types"
	"lawfirm-go-backend/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	clientStore types.ClientStore
	userStore   types.UserStore
}

func NewHandler(clientStore types.ClientStore, userStore types.UserStore) *Handler {
	return &Handler{
		clientStore: clientStore,
		userStore:   userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/client/{clientID}", auth.WithJWTAuth(h.handleGetClient, h.userStore, true)).Methods("GET")
	router.HandleFunc("/client", auth.WithJWTAuth(h.handleGetAllClients, h.userStore, true)).Methods("GET")
	router.HandleFunc("/client", auth.WithJWTAuth(h.handleCreateClient, h.userStore, true)).Methods("POST")
	router.HandleFunc("/client", auth.WithJWTAuth(h.handleUpdateClient, h.userStore, true)).Methods("PUT")
	router.HandleFunc("/client/{clientID}", auth.WithJWTAuth(h.handleDeleteClient, h.userStore, true)).Methods("DELETE")
}

func (h *Handler) handleGetAllClients(w http.ResponseWriter, r *http.Request) {
	clients, err := h.clientStore.GetAllClients()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, clients)
}

func (h *Handler) handleGetClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str := vars["clientID"]
	clientID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Client ID"))
		return
	}

	c, err := h.clientStore.GetClientByID(clientID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, c)
}

func (h *Handler) handleCreateClient(w http.ResponseWriter, r *http.Request) {
	// Get JSON Payload
	var payload types.ClientPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the Payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", errors))
		return
	}

	// Validate that the client doesn't exist
	_, err := h.clientStore.GetClientByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Client with email %s already exists", payload.Email))
		return
	}

	// If the client doesn't exist, create new client
	err = h.clientStore.CreateClient(types.Client{
		Name:        payload.Name,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleUpdateClient(w http.ResponseWriter, r *http.Request) {
	// Get JSON Payload
	var payload types.ClientPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the Payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", errors))
		return
	}

	// Validate that the client does exist
	c, err := h.clientStore.GetClientByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Client with email %s does not exist", payload.Email))
		return
	}

	// If the client does exist, update the client
	err = h.clientStore.UpdateClient(types.Client{
		ID:          c.ID,
		Name:        payload.Name,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleDeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str := vars["clientID"]
	clientID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Client ID"))
		return
	}

	// Validate that the client does exist
	c, err := h.clientStore.GetClientByID(clientID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Client with id %s does not exist", clientID))
		return
	}

	// If the client does exist, delete the client
	err = h.clientStore.DeleteClient(c.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
