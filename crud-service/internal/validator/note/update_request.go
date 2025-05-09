package validator

type UpdateNoteRequest struct {
	Title   string `json:"title" binding:"omitempty,min=2,max=100"`
	Content string `json:"content" binding:"omitempty,min=2,max=1000"`
}
