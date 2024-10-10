package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rfulgencio3/go-personal-library/internal/domain"
	"github.com/rfulgencio3/go-personal-library/internal/usecase"
	"github.com/rfulgencio3/go-personal-library/internal/validator"
)

type BookHandler struct {
	bookUseCase usecase.BookUseCase
}

func NewBookHandler(bu usecase.BookUseCase) *BookHandler {
	return &BookHandler{
		bookUseCase: bu,
	}
}

func (h *BookHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/books", h.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", h.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods("DELETE")
	router.HandleFunc("/books", h.GetAllBooks).Methods("GET")
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the library
// @Tags books
// @Accept json
// @Produce json
// @Param book body domain.Book true "Book to add"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books [post]
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Chamar o caso de uso para criar o livro
	if err := h.bookUseCase.CreateBook(&book); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to create book")
		return
	}

	// Retornar o livro criado, incluindo o ID gerado
	h.respondWithJSON(w, http.StatusCreated, SuccessResponse{Data: book})
}

// GetBookByID godoc
// @Summary Get a book by ID
// @Description Retrieve a book from the library by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} SuccessResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/{id} [get]
func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	book, err := h.bookUseCase.GetBookByID(id)
	if err != nil {
		h.respondWithError(w, http.StatusNotFound, "Book not found")
		return
	}
	h.respondWithJSON(w, http.StatusOK, SuccessResponse{Data: book})
}

// UpdateBook godoc
// @Summary Update a book by ID
// @Description Update a book in the library by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body domain.Book true "Updated book data"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	book.ID = id
	if err := h.bookUseCase.UpdateBook(&book); err != nil {
		if errors.Is(err, validator.ErrInvalidBookData) {
			h.respondWithError(w, http.StatusBadRequest, err.Error())
		} else {
			h.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}
	h.respondWithJSON(w, http.StatusOK, SuccessResponse{Data: book})
}

// DeleteBook godoc
// @Summary Delete a book by ID
// @Description Remove a book from the library by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 204
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := h.bookUseCase.DeleteBook(id); err != nil {
		h.respondWithError(w, http.StatusNotFound, "Book not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Retrieve all books from the library
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 500 {object} ErrorResponse
// @Router /books [get]
func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookUseCase.GetAllBooks()
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	h.respondWithJSON(w, http.StatusOK, SuccessResponse{Data: books})
}

// Helper functions
func (h *BookHandler) respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func (h *BookHandler) respondWithError(w http.ResponseWriter, statusCode int, message string) {
	h.respondWithJSON(w, statusCode, ErrorResponse{Message: message})
}
