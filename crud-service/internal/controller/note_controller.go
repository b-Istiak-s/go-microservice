package controller

import (
	"crud-service/internal/model"
	"crud-service/internal/repository"
	"crud-service/internal/util/response"
	validator "crud-service/internal/validator/note"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	noteRepository repository.NoteRepository
}

func NewNoteController(noteRepository repository.NoteRepository) *NoteController {
	return &NoteController{
		noteRepository: noteRepository,
	}
}
func (noteController *NoteController) CreateNote(c *gin.Context) {
	var req validator.CreateNoteRequest

	// Bind and validate request
	ok, errs := validator.BindAndValidateNote(c, &req)
	if !ok {
		response.Success(c, http.StatusUnprocessableEntity, "Validation Error", errs)
		return
	}

	// Get userID from context
	userIDInterface, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "User ID not found in context")
		return
	}

	// Gin context stores values as interface{}, so you need to cast it
	userID, ok := userIDInterface.(uint)
	if !ok {
		response.Error(c, http.StatusInternalServerError, "Invalid user ID type")
		return
	}

	note := model.Note{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := noteController.noteRepository.Create(&note); err != nil {
		response.Error(c, http.StatusBadRequest, "Error creating note", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Note created successfully", note)
}
func (nc *NoteController) GetAllNotes(c *gin.Context) {
	// Implementation for getting all notes
}
func (nc *NoteController) GetNoteByID(c *gin.Context) {
	// Implementation for getting a note by ID
}
func (nc *NoteController) UpdateNote(c *gin.Context) {
	// Implementation for updating a note
}
func (nc *NoteController) DeleteNote(c *gin.Context) {
	// Implementation for deleting a note
}
