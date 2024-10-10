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
	router.HandleFunc("/read_books", h.GetAllReadBooks).Methods("GET")
	router.HandleFunc("/read_books/{id}", h.UpdateReadBook).Methods("PUT")
	router.HandleFunc("/read_books/{id}", h.DeleteReadBook).Methods("DELETE")
	router.HandleFunc("/read_books/{id}/comments", h.AddCommentToReadBook).Methods("POST")
}

// CreateReadBook godoc
// @Summary Create a new read book
// @Description Add a new read book record
// @Tags read_books
// @Accept json
// @Produce json
// @Param read_book body domain.ReadBook true "Read Book to add (no ID)"
// @Success 201 {object} domain.ReadBook
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /read_books [post]
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

// GetReadBookByID godoc
// @Summary Get a read book by ID
// @Description Get a read book record by its ID
// @Tags read_books
// @Accept json
// @Produce json
// @Param id path string true "Read Book ID"
// @Success 200 {object} domain.ReadBook
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /read_books/{id} [get]
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

// GetAllReadBooks godoc
// @Summary Get all read books
// @Description Get all read book records
// @Tags read_books
// @Accept json
// @Produce json
// @Success 200 {array} domain.ReadBook
// @Failure 500 {object} ErrorResponse
// @Router /read_books [get]
func (h *ReadBookHandler) GetAllReadBooks(w http.ResponseWriter, r *http.Request) {
	readBooks, err := h.usecase.GetAllReadBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(readBooks)
}

// UpdateReadBook godoc
// @Summary Update a read book by ID
// @Description Update a read book record by its ID
// @Tags read_books
// @Accept json
// @Produce json
// @Param id path string true "Read Book ID"
// @Param read_book body domain.ReadBook true "Updated Read Book data"
// @Success 200 {object} domain.ReadBook
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /read_books/{id} [put]
func (h *ReadBookHandler) UpdateReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var readBook domain.ReadBook
	if err := json.NewDecoder(r.Body).Decode(&readBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	readBook.ID = id
	if err := h.usecase.UpdateReadBook(&readBook); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(readBook)
}

// DeleteReadBook godoc
// @Summary Delete a read book by ID
// @Description Delete a read book record by its ID
// @Tags read_books
// @Accept json
// @Produce json
// @Param id path string true "Read Book ID"
// @Success 204
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /read_books/{id} [delete]
func (h *ReadBookHandler) DeleteReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.usecase.DeleteReadBook(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AddCommentToReadBook godoc
// @Summary Add a comment to a read book
// @Description Add a comment to the read book's comments list
// @Tags read_books
// @Accept json
// @Produce json
// @Param id path string true "Read Book ID"
// @Param comment body string true "Comment to add"
// @Success 200 {string} string "Comment added successfully"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /read_books/{id}/comments [post]
func (h *ReadBookHandler) AddCommentToReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["book_id"]

	var comment struct {
		Comment string `json:"comment"`
	}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.usecase.AddCommentToReadBook(bookId, comment.Comment); err != nil {
		if err.Error() == "read book not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Comment added successfully")
}
