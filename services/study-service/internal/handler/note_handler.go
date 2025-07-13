package handler

import (
	"net/http"
	"study-service/internal/service"

	"github.com/gin-gonic/gin"
)

// createNoteRequest, CreateNote handler'ı için gelen isteğin gövdesini temsil eder.
type createNoteRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	CourseCode  string `json:"courseCode" binding:"required"`
}

// NoteHandler, notlarla ilgili HTTP isteklerini yönetir.
type NoteHandler struct {
	noteService service.NoteService
}

// NewNoteHandler, yeni bir NoteHandler örneği oluşturur.
func NewNoteHandler(s service.NoteService) *NoteHandler {
	return &NoteHandler{
		noteService: s,
	}
}

// CreateNote, yeni bir not oluşturmak için kullanılan gin handler'ıdır.
func (h *NoteHandler) CreateNote(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-ID header is required"})
		return
	}

	var req createNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := h.noteService.Create(c.Request.Context(), userID, req.Title, req.Description, req.CourseCode)
	if err != nil {
		// Burada daha detaylı hata yönetimi yapılabilir (örn: service katmanından gelen hatanın türüne göre farklı status kodları dönmek)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, note)
}
