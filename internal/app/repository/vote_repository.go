package repository

import (
	"WebApp/internal/app/auth"
	"WebApp/internal/app/ds"
	"errors"
	"time"

	"gorm.io/gorm"
)

/* ==========================================
   DRAFT VOTE (КОРЗИНА)
========================================== */

func (r *Repository) GetOrCreateDraftVote() (
	*ds.Vote,
	error,
) {

	userID := auth.GetCurrentUserID()

	var vote ds.Vote

	err := r.db.
		Where("creator_id = ?", userID).
		Where("status = ?", ds.VoteDraft).
		First(&vote).Error

	// уже существует
	if err == nil {
		return &vote, nil
	}

	// ошибка БД
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// создаём новую корзину
	vote = ds.Vote{
		CreatorID: userID,
		Status:    ds.VoteDraft,

	}

	err = r.db.Create(&vote).Error
	if err != nil {
		return nil, err
	}

	return &vote, nil
}

/* ==========================================
   CART ICON
========================================== */

func (r *Repository) GetCartInfo() (
	map[string]interface{},
	error,
) {

	vote, err := r.GetOrCreateDraftVote()
	if err != nil {
		return nil, err
	}

	var count int64

	err = r.db.
		Model(&ds.VoteItem{}).
		Where("vote_id = ?", vote.ID).
		Count(&count).Error

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"vote_id": vote.ID,
		"count":   count,
	}, nil
}

/* ==========================================
   GET LIST
========================================== */

func (r *Repository) GetVotes(
	status string,
	dateFrom string,
	dateTo string,
) ([]ds.Vote, error) {

	var votes []ds.Vote

	query := r.db.
		Model(&ds.Vote{}).
		Preload("Creator").
		Preload("Moderator").
		Where("status != ?", ds.VoteDeleted)

	if status != "" {
		query = query.Where(
			"status = ?",
			status,
		)
	}

	if dateFrom != "" {
		from, err := time.Parse(
			"2006-01-02",
			dateFrom,
		)

		if err == nil {
			query = query.Where(
				"created_at >= ?",
				from,
			)
		}
	}

	if dateTo != "" {
		to, err := time.Parse(
			"2006-01-02",
			dateTo,
		)

		if err == nil {
			query = query.Where(
				"created_at <= ?",
				to,
			)
		}
	}

	err := query.
		Order("created_at DESC").
		Find(&votes).Error

	return votes, err
}

/* ==========================================
   GET BY ID
========================================== */

func (r *Repository) GetVoteByID(
	id uint,
) (*ds.Vote, error) {

	var vote ds.Vote

	err := r.db.
		Preload("Creator").
		Preload("Moderator").
		Preload("VoteItems").
		Preload("VoteItems.Event").
		First(&vote, id).Error

	if err != nil {
		return nil, err
	}

	return &vote, nil
}

/* ==========================================
   FORM VOTE
========================================== */

func (r *Repository) FormVote(
	id uint,
) error {

	vote, err := r.GetVoteByID(id)
	if err != nil {
		return err
	}

	if vote.Status != ds.VoteDraft {
		return errors.New(
			"only draft can be formed",
		)
	}

	var items []ds.VoteItem

	err = r.db.
		Where("vote_id = ?", vote.ID).
		Find(&items).Error

	if err != nil {
		return err
	}

	if len(items) == 0 {
		return errors.New(
			"vote is empty",
		)
	}

	result := 0

	for _, item := range items {
		result += item.Value
	}

	now := time.Now()

	vote.Result = result
	vote.Status = ds.VoteFormed
	vote.FormedAt = &now

	return r.db.Save(vote).Error
}

/* ==========================================
   COMPLETE
========================================== */

func (r *Repository) CompleteVote(
	id uint,
) error {

	vote, err := r.GetVoteByID(id)
	if err != nil {
		return err
	}

	if vote.Status != ds.VoteFormed {
		return errors.New(
			"only formed vote can be completed",
		)
	}

	now := time.Now()
	moderatorID := auth.GetCurrentUserID()

	vote.Status = ds.VoteCompleted
	vote.CompletedAt = &now
	vote.ModeratorID = &moderatorID

	return r.db.Save(vote).Error
}

/* ==========================================
   REJECT
========================================== */

func (r *Repository) RejectVote(
	id uint,
) error {

	vote, err := r.GetVoteByID(id)
	if err != nil {
		return err
	}

	if vote.Status != ds.VoteFormed {
		return errors.New(
			"only formed vote can be rejected",
		)
	}

	now := time.Now()
	moderatorID := auth.GetCurrentUserID()

	vote.Status = ds.VoteRejected
	vote.CompletedAt = &now
	vote.ModeratorID = &moderatorID

	return r.db.Save(vote).Error
}

/* ==========================================
   DELETE
========================================== */

func (r *Repository) DeleteVote(
	id uint,
) error {

	vote, err := r.GetVoteByID(id)
	if err != nil {
		return err
	}

	if vote.Status != ds.VoteDraft {
		return errors.New(
			"only draft vote can be deleted",
		)
	}

	vote.Status = ds.VoteDeleted

	return r.db.Save(vote).Error
}
