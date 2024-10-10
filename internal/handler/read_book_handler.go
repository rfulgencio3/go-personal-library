package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rfulgencio3/go-personal-library/internal/domain"
	"github.com/rfulgencio3/go-personal-library/internal/usecase"
)

type ReadBookHandler struct {
	usecase usecase.ReadBookUseCase
}

func NewReadBookHandler(uc usecase.ReadBookUseCase) *ReadBookHandler {
	return &ReadBookHandler{usecase: uc}
}

func (h *ReadBookHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/read_books", h.CreateReadBook).Methods("POST")
	router.HandleFunc("/read_books/{id}", h.GetReadBookByID).Methods("GET")
	// Adicione outros endpoints para Update e Delete...
}

func (h *ReadBookHandler) CreateReadBook(w http.ResponseWriter, r *http.Request) {
	var readBook domain.ReadBook
	if err := json.NewDecoder(r.Body).Decode(&readBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.usecase.CreateReadBook(&readBook); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(readBook)
}

func (h *ReadBookHandler) GetReadBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	readBook, err := h.usecase.GetReadBookByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(readBook)
}
