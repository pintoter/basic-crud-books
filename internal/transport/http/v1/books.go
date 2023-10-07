package v1

import (
	"books-app/internal/model"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initBooksRoutes(api *gin.RouterGroup) {
	api.POST("/create", h.createBook)
	api.GET("/", h.getAllBooks)
	api.GET("/:id", h.getBookById)
	api.GET("/getByAuthor", h.getBooksByAuthor)
	api.PUT("/update_book", h.updateBook)
	api.DELETE("/delete_by_id", h.deleteBookByID)
	api.DELETE("/delete_by_title", h.deleteBookByTitle)
}

type createBookInput struct {
	Author          string  `json:"author"`
	Title           string  `json:"title"`
	Rating          float64 `json:"rating"`
	PublicationYear int64   `json:"year"`
}

type createBookResponse struct {
	ID int `json:"id"`
}

func (h *Handler) createBook(c *gin.Context) {
	var input createBookInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, model.ErrInvalidData.Error())
		return
	}

	id, err := h.service.Create(c, model.Book{Author: input.Author, Title: input.Title, Rating: input.Rating, PublicationYear: input.PublicationYear})
	if err != nil {
		if errors.Is(err, model.ErrBookAlreadyExists) {
			newErrorResponse(c, http.StatusConflict, err.Error())
		} else if errors.Is(err, model.ErrInvalidData) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	newResponse(c, http.StatusCreated, createBookResponse{ID: id})
}

type getAllBooksResponse struct {
	Books []model.Book `json:"books"`
}

func (h *Handler) getAllBooks(c *gin.Context) {
	books, err := h.service.GetAll(c)
	if err != nil {
		if errors.Is(err, model.ErrBooksNotFound) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
	}

	newResponse(c, http.StatusOK, getAllBooksResponse{Books: books})
}

type getBookResponse struct {
	Book model.Book `json:"book"`
}

func (h *Handler) getBookById(c *gin.Context) {
	paramId := strings.Trim(c.Param("id"), "/")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid parameter (id)")
	}

	book, err := h.service.GetByID(c, id)
	if err != nil {
		if errors.Is(err, model.ErrBookNotFound) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
	}

	newResponse(c, http.StatusOK, getBookResponse{Book: book})
}

type getBooksResponse struct {
	Books []model.Book `json:"books"`
}

type getByAuthorInput struct {
	Author string `json:"author" binding:"required"`
}

func (h *Handler) getBooksByAuthor(c *gin.Context) {
	var input getByAuthorInput
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid author")
		return
	}

	books, err := h.service.GetByAuthor(c, input.Author)
	if err != nil {
		if errors.Is(err, model.ErrBooksNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error()) // я сделал statusnotfound был statusbadrequest
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
	}

	newResponse(c, http.StatusOK, getBooksResponse{Books: books})
}

type updateBookInput struct {
	Author string  `json:"author" binding:"required"`
	Title  string  `json:"title" binding:"required"`
	Rating float64 `json:"rating" binding:"required"`
}

func (h *Handler) updateBook(c *gin.Context) {
	var input updateBookInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input parameters")
		return
	}

	fmt.Println(input.Author, input.Rating, input.Title)

	err := h.service.UpdateBook(c, input.Rating, input.Author, input.Title)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	newResponse(c, http.StatusNoContent, nil)
}

type deleteBookByIDInput struct {
	ID int `json:"id"`
}

func (h *Handler) deleteBookByID(c *gin.Context) {
	var input deleteBookByIDInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid parameter (id)")
		return
	}

	err := h.service.DeleteByID(c, input.ID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid something...")
	}

	newResponse(c, http.StatusNoContent, nil)
}

type deleteBookByTitleInput struct {
	Title string `json:"title"`
}

func (h *Handler) deleteBookByTitle(c *gin.Context) {
	var input deleteBookByTitleInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid parameter (id)")
		return
	}

	err := h.service.DeleteByTitle(c, input.Title)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid something...")
	}

	newResponse(c, http.StatusNoContent, nil)
}
