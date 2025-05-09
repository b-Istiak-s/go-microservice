package controller

import (
	"crud-service/internal/model"
	"crud-service/internal/repository"
	"crud-service/internal/util/response"
	validator "crud-service/internal/validator/note"
	"net/http"
	"strconv"

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

	userID, err := GetUserID(c)
	if !err {
		return // Error response already sent in GetUserID
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
	userID, errBool := GetUserID(c)
	if !errBool {
		return // Error response already sent in GetUserID
	}
	notes, err := nc.noteRepository.GetAll(userID)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Error fetching notes")
		return
	}

	if len(notes) == 0 {
		response.Success(c, http.StatusOK, "No notes found", nil)
		return
	}

	response.Success(c, http.StatusOK, "Notes fetched successfully", notes)
}
func (nc *NoteController) UpdateNote(c *gin.Context) {
	var req validator.UpdateNoteRequest
	// Bind and validate request
	ok, errs := validator.BindAndValidateNote(c, &req)
	if !ok {
		response.Success(c, http.StatusUnprocessableEntity, "Validation Error", errs)
		return
	}

	userID, errBool := GetUserID(c)
	if !errBool {
		return // Error response already sent in GetUserID
	}

	// Fetch note ID from URL parameters
	noteIDParam := c.Param("id")
	noteID, err := strconv.ParseUint(noteIDParam, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid note ID")
		return
	}

	note, err := nc.noteRepository.GetByID(uint(noteID))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Note not found", err.Error())
		return
	}
	if note.UserID != userID {
		response.Error(c, http.StatusForbidden, "You do not have permission to update this note", nil)
		return
	}

	if req.Title != "" {
		note.Title = req.Title
	}
	if req.Content != "" {
		note.Content = req.Content
	}

	if err := nc.noteRepository.Update(note); err != nil {
		response.Error(c, http.StatusInternalServerError, "Error updating note")
		return
	}

	response.Success(c, http.StatusOK, "Note updated successfully", note)
}

func (nc *NoteController) DeleteNote(c *gin.Context) {
	// Implementation for deleting a note
}

func GetUserID(c *gin.Context) (uint, bool) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "User ID not found in context")
		return 0, false
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		response.Error(c, http.StatusInternalServerError, "Invalid user ID type")
		return 0, false
	}

	return userID, true
}
