package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCart(
	ctx *gin.Context,
) {

	data, err := h.Repository.
		GetCartInfo()

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		data,
	)
}

func (h *Handler) GetVotes(
	ctx *gin.Context,
) {

	status := ctx.Query("status")
	dateFrom := ctx.Query("date_from")
	dateTo := ctx.Query("date_to")

	votes, err := h.Repository.
		GetVotes(
			status,
			dateFrom,
			dateTo,
		)

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		votes,
	)
}

func (h *Handler) GetVoteByID(
	ctx *gin.Context,
) {

	id, _ := strconv.Atoi(
		ctx.Param("id"),
	)

	vote, err := h.Repository.
		GetVoteByID(
			uint(id),
		)

	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			gin.H{"error": "vote not found"},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		vote,
	)
}

func (h *Handler) FormVote(
	ctx *gin.Context,
) {

	id, _ := strconv.Atoi(
		ctx.Param("id"),
	)

	err := h.Repository.FormVote(
		uint(id),
	)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "vote formed",
		},
	)
}

func (h *Handler) CompleteVote(
	ctx *gin.Context,
) {

	id, _ := strconv.Atoi(
		ctx.Param("id"),
	)

	err := h.Repository.
		CompleteVote(uint(id))

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "completed",
		},
	)
}

func (h *Handler) RejectVote(
	ctx *gin.Context,
) {

	id, _ := strconv.Atoi(
		ctx.Param("id"),
	)

	err := h.Repository.
		RejectVote(
			uint(id),
		)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "rejected",
		},
	)
}

func (h *Handler) DeleteVote(
	ctx *gin.Context,
) {

	id, _ := strconv.Atoi(
		ctx.Param("id"),
	)

	err := h.Repository.
		DeleteVote(
			uint(id),
		)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.Status(http.StatusNoContent)
}
