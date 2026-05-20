package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VoteItemInput struct {
	EventID uint `json:"event_id"`
	Priority int `json:"priority"`
	Value int `json:"value"`
}

func (h *Handler) AddVoteItem(
	ctx *gin.Context,
) {

	var input VoteItemInput

	if err := ctx.ShouldBindJSON(
		&input,
	); err != nil {

		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	err := h.Repository.
		AddEventToVote(
			input.EventID,
			input.Priority,
			input.Value,
		)

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"message": "added",
		},
	)
}

func (h *Handler) UpdateVoteItem(
	ctx *gin.Context,
) {

	eventID, _ := strconv.Atoi(
		ctx.Param("event_id"),
	)

	var input VoteItemInput

	ctx.ShouldBindJSON(&input)

	err := h.Repository.
		UpdateVoteItem(
			uint(eventID),
			input.Priority,
			input.Value,
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
		gin.H{"message": "updated"},
	)
}

func (h *Handler) DeleteVoteItem(
	ctx *gin.Context,
) {

	eventID, _ := strconv.Atoi(
		ctx.Param("event_id"),
	)

	err := h.Repository.
		DeleteVoteItem(
			uint(eventID),
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