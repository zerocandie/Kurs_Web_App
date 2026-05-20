package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// VoteItemInput модель связи услуги с заявкой
// @Description Данные для many-to-many связи
type VoteItemInput struct {
	EventID  uint `json:"event_id" example:"1" binding:"required"`
	Priority int  `json:"priority" example:"1" binding:"min=1"`
	Value    int  `json:"value" example:"100" binding:"min=0"`
}

// AddVoteItem godoc
// @Summary      Добавить услугу в заявку
// @Tags         VoteItems
// @Accept       json
// @Produce      json
// @Param        request  body  VoteItemInput  true  "Данные связи"
// @Success      201  {object}  map[string]string  "Добавлено"
// @Failure      400  {object}  map[string]string  "Ошибка валидации"
// @Failure      500  {object}  map[string]string  "Ошибка сервера"
// @Security     BearerAuth
// @Router       /vote-items [post]
func (h *Handler) AddVoteItem(ctx *gin.Context) {
	var input VoteItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Repository.AddEventToVote(input.EventID, input.Priority, input.Value)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "added"})
}

// UpdateVoteItem godoc
// @Summary      Изменить услугу в заявке
// @Tags         VoteItems
// @Accept       json
// @Produce      json
// @Param        event_id  path  uint  true  "ID связи"
// @Param        request   body  VoteItemInput  true  "Новые данные"
// @Success      200  {object}  map[string]string  "Обновлено"
// @Failure      400  {object}  map[string]string  "Ошибка"
// @Failure      404  {object}  map[string]string  "Не найдено"
// @Security     BearerAuth
// @Router       /vote-items/{event_id} [put]
func (h *Handler) UpdateVoteItem(ctx *gin.Context) {
	idStr := ctx.Param("event_id")
	eventID, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	if eventID < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be positive"})
		return
	}
	var input VoteItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.Repository.UpdateVoteItem(uint(eventID), input.Priority, input.Value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// DeleteVoteItem godoc
// @Summary      Удалить услугу из заявки
// @Tags         VoteItems
// @Produce      json
// @Param        event_id  path  uint  true  "ID связи"
// @Success      204  {string}  string  "No Content"
// @Failure      400  {object}  map[string]string  "Ошибка"
// @Failure      404  {object}  map[string]string  "Не найдено"
// @Security     BearerAuth
// @Router       /vote-items/{event_id} [delete]
func (h *Handler) DeleteVoteItem(ctx *gin.Context) {
	idStr := ctx.Param("event_id")
	eventID, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	if eventID < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be positive"})
		return
	}
	err = h.Repository.DeleteVoteItem(uint(eventID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
