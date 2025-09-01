package handler

import (
	"net/http"
	"study-service/internal/domain"
	"study-service/internal/service"

	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	noteService service.NoteService
}

func NewNoteHandler (service service.NoteService) *NoteHandler {
	return &NoteHandler{noteService: service}
}

func (h *NoteHandler) RegisterRoutes(router *gin.RouterGroup) {
	notes := router.Group("/notes")
	{
		notes.POST("", h.CreateNote)
		notes.GET("/:id", h.GetNoteById)
		notes.PUT("/:id", h.UpdateNote)
		notes.DELETE("/:id", h.DeleteNote)
	}
}

func (h *NoteHandler) CreateNote(c *gin.Context) {
	var note domain.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri formatı"})
		return
	}
	createdNote, err := h.noteService.Create(c.Request.Context(), note)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, createdNote)
}

func (h *NoteHandler) GetNoteById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz id"})
		return
	}
	note, err := h.noteService.GetById(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, note)
}

func (h *NoteHandler) UpdateNote(c *gin.Context) {
	noteId := c.Param("id")
	if noteId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz not ID"})
		return
	}
	var partialUpdate domain.Note
	if err := c.ShouldBindJSON(&partialUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri formatı"})
		return
	}
	partialUpdate.ID = noteId 

	updatedNote, err := h.noteService.Update(c.Request.Context(), partialUpdate)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, updatedNote)
}

func (h *NoteHandler) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz id"})
		return
	}
	err := h.noteService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}