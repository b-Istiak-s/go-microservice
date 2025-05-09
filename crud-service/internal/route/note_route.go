package route

import (
	"crud-service/internal/controller"
	"crud-service/internal/db"
	"crud-service/internal/middleware"
	"crud-service/internal/repository"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(router *gin.Engine) {
	// Set up repository
	noteRepo := repository.NewNoteRepository(db.DB)

	// Set up controller
	noteController := controller.NewNoteController(noteRepo)

	// Group your routes under /api/notes
	noteRoutes := router.Group("/api/notes", middleware.AuthMiddleware())
	{
		noteRoutes.POST("/", noteController.CreateNote)
		noteRoutes.GET("/", noteController.GetAllNotes)
		noteRoutes.PUT("/:id", noteController.UpdateNote)
	}
}
