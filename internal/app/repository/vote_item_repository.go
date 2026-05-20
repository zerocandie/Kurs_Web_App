package repository

import "WebApp/internal/app/ds"

func (r *Repository) AddEventToVote(
	eventID uint,
	priority int,
	value int,
) error {

	vote, err := r.GetOrCreateDraftVote()
	if err != nil {
		return err
	}

	item := ds.VoteItem{
		VoteID: vote.ID,
		EventID: eventID,
		Priority: priority,
		Value: value,
	}

	return r.db.Create(&item).Error
}

func (r *Repository) UpdateVoteItem(
	eventID uint,
	priority int,
	value int,
) error {

	vote, err := r.GetOrCreateDraftVote()
	if err != nil {
		return err
	}

	return r.db.Model(
		&ds.VoteItem{},
	).
		Where(
			"vote_id = ?",
			vote.ID,
		).
		Where(
			"event_id = ?",
			eventID,
		).
		Updates(map[string]interface{}{
			"priority": priority,
			"value": value,
		}).Error
}

func (r *Repository) DeleteVoteItem(
	eventID uint,
) error {

	vote, err := r.GetOrCreateDraftVote()
	if err != nil {
		return err
	}

	return r.db.
		Where(
			"vote_id = ?",
			vote.ID,
		).
		Where(
			"event_id = ?",
			eventID,
		).
		Delete(&ds.VoteItem{}).
		Error
}