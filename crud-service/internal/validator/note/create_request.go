package validator

type CreateNoteRequest struct {
	Title   string `json:"title" binding:"required,min=2,max=100"`
	Content string `json:"content" binding:"required,min=2,max=1000"`
}
