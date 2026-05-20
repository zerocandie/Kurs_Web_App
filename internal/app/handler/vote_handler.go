package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCart godoc
// @Summary      Иконка корзины
// @Tags         Votes
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "Данные корзины"
// @Failure      500  {object}  map[string]string  "Ошибка сервера"
// @Security     BearerAuth
// @Router       /votes/cart [get]
func (h *Handler) GetCart(ctx *gin.Context) {
	data, err := h.Repository.GetCartInfo()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// GetVotes godoc
// @Summary      Список заявок с фильтрацией
// @Tags         Votes
// @Produce      json
// @Param        status     query  string  false  "Статус"
// @Param        date_from  query  string  false  "Дата от"
// @Param        date_to    query  string  false  "Дата до"
// @Success      200  {object}  map[string]interface{}  "Список заявок"
// @Failure      500  {object}  map[string]string  "Ошибка сервера"
// @Security     BearerAuth
// @Router       /votes [get]
func (h *Handler) GetVotes(ctx *gin.Context) {
	status := ctx.Query("status")
	dateFrom := ctx.Query("date_from")
	dateTo := ctx.Query("date_to")
	votes, err := h.Repository.GetVotes(status, dateFrom, dateTo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, votes)
}

// GetVoteByID godoc
// @Summary      Заявка по ID
// @Tags         Votes
// @Produce      json
// @Param        id  path  uint  true  "ID заявки"
// @Success      200  {object}  map[string]interface{}  "Данные заявки"
// @Failure      400  {object}  map[string]string  "Невалидный ID"
// @Failure      404  {object}  map[string]string  "Не найдено"
// @Failure      500  {object}  map[string]string  "Ошибка сервера"
// @Security     BearerAuth
// @Router       /votes/{id} [get]
func (h *Handler) GetVoteByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	if idInt < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be positive"})
		return
	}
	vote, err := h.Repository.GetVoteByID(uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "vote not found"})
		return
	}
	ctx.JSON(http.StatusOK, vote)
}

// FormVote godoc
// @Summary      Сформировать заявку
// @Tags         Votes
// @Produce      json
// @Param        id  path  uint  true  "ID заявки"
// @Success      200  {object}  map[string]string  "Заявка сформирована"
// @Failure      400  {object}  map[string]string  "Ошибка"
// @Failure      404  {object}  map[string]string  "Не найдено"
// @Security     BearerAuth
// @Router       /votes/{id}/form [put]
func (h *Handler) FormVote(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	if idInt < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be positive"})
		return
	}
	err = h.Repository.FormVote(uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "vote formed"})
}

// CompleteVote godoc
// @Summary      Завершить заявку (модератор)
// @Tags         Votes
// @Produce      json
// @Param        id  path  uint  true  "ID заявки"
// @Success      200  {object}  map[string]string  "Заявка завершена"
// @Failure      400  {object}  map[string]string  "Ошибка"
// @Failure      403  {object}  map[string]string  "Доступ запрещён"
// @Failure      404  {object}  map[string]string  "Не найдено"
// @Security     BearerAuth
// @Router       /votes/{id}/complete [put]
func (h *Handler) CompleteVote(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	if idInt < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be positive"})
		return
	}
	err = h.Repository.CompleteVote(uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "completed"})
}

// RejectVote godoc
// @Summary      Отклонить заявку (модератор)
// @Tags         Votes
// @Produce      json
// @Param        id  path  uint  true  "ID заявки"
// @Success      200  {object}  map[string]string  "Заявка отклонена"
// @Failure      400  {object}  map[string]string  "Ошибка"
// @Failure      403  {object}  map[string]string  "Доступ запрещён"
// @Failure      404  {object}  map[string]string  "Не найдено"
// @Security     BearerAuth
// @Router       /votes/{id}/reject [put]
func (h *Handler) RejectVote(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	if idInt < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be positive"})
		return
	}
	err = h.Repository.RejectVote(uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "rejected"})
}

// DeleteVote godoc
// @Summary      Удалить заявку
// @Tags         Votes
// @Produce      json
// @Param        id  path  uint  true  "ID заявки"
// @Success      204  {string}  string  "No Content"
// @Failure      400  {object}  map[string]string  "Ошибка"
// @Failure      404  {object}  map[string]string  "Не найдено"
// @Security     BearerAuth
// @Router       /votes/{id} [delete]
func (h *Handler) DeleteVote(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	if idInt < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be positive"})
		return
	}
	err = h.Repository.DeleteVote(uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
