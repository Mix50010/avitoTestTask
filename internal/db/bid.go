package db

import (
	"context"
	"tenderService/internal/models"
	"time"
)

// CreateBid creates a new bid based on the provided request.
//
// models.CreateBidRequest contains the necessary information for bid creation.
// Returns the created bid ID, creation time, and an error if any.
func CreateBid(req models.CreateBidRequest) (string, time.Time, error) {
	query := `
	INSERT INTO public.bids(
	name, description, status, tender_id, author_type, author_id, version)
	VALUES ($1, $2, 'Created', $3, $4, $5, 1)
	RETURNING id, created_at;
	`

	// Выполняем запрос
	var bidID string
	var createdAt time.Time
	err := db.QueryRow(
		context.Background(),
		query,
		req.Name,
		req.Description,
		req.TenderId,
		req.AuthorType,
		req.AuthorId,
	).Scan(&bidID, &createdAt)
	if err != nil {
		return "", time.Now(), err
	}
	return bidID, time.Now(), nil
}

func GetBidsByAuthorID(authorID string, limit, offset int) ([]models.Bid, error) {
	var bids []models.Bid
	rows, err := db.Query(
		context.Background(),
		`SELECT id, name, description, status, tender_id, author_type, author_id, version, created_at
		FROM public.bids
		WHERE author_id = $1
		ORDER BY name ASC
		LIMIT $2 OFFSET $3`,
		authorID,
		limit,
		offset)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bid models.Bid
		err := rows.Scan(&bid.Id, &bid.Name, &bid.Description, &bid.Status, &bid.TenderId, &bid.AuthorType, &bid.AuthorId, &bid.Version, &bid.CreatedAt)
		if err != nil {
			return nil, err
		}
		bids = append(bids, bid)
	}
	return bids, nil
}

func GetBidsByTenderID(tenderID string, limit, offset int) ([]models.Bid, error) {
	var bids []models.Bid
	rows, err := db.Query(
		context.Background(),
		`SELECT id, name, description, status, tender_id, author_type, author_id, version, created_at
		FROM public.bids
		WHERE tender_id = $1
		ORDER BY name ASC
		LIMIT $2 OFFSET $3`,
		tenderID,
		limit,
		offset)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bid models.Bid
		err := rows.Scan(&bid.Id, &bid.Name, &bid.Description, &bid.Status, &bid.TenderId, &bid.AuthorType, &bid.AuthorId, &bid.Version, &bid.CreatedAt)
		if err != nil {
			return nil, err
		}
		bids = append(bids, bid)
	}
	return bids, nil
}

// GetBidStatusByID возвращает статус предложения по его идентификатору.
//
// bidID - идентификатор предложения.
// Возвращает статус предложения и ошибку, если произошла ошибка.
func GetBidStatusByID(bidID string) (string, error) {
	var status string
	err := db.QueryRow(context.Background(),
		`SELECT status
		FROM bids
		WHERE id = $1;`,
		bidID).Scan(&status)
	if err != nil {
		return "", err
	}
	return status, nil
}


// SetBidStatusByID updates the status of a bid with the given ID.
//
// Parameters:
// - bidID: the ID of the bid to update.
// - status: the new status to set.
//
// Returns:
// - error: an error if the update operation fails.
func SetBidStatusByID(bidID string, status string) error {
	_, err := db.Exec(context.Background(),
		`UPDATE bids
		SET status = $1
		WHERE id = $2;`,
		status,
		bidID)
	return err
}

// GetBidByID возвращает предложение по его идентификатору.
//
// bidID - идентификатор предложения.
// Возвращает предложение и ошибку, если произошла ошибка.
func GetBidByID(bidID string) (models.Bid, error) {
	var bid models.Bid
	err := db.QueryRow(context.Background(),
		`SELECT id, name, description, status, tender_id, author_type, author_id, version, created_at
		FROM public.bids
		WHERE id = $1;`,
		bidID).Scan(&bid.Id, &bid.Name, &bid.Description, &bid.Status, &bid.TenderId, &bid.AuthorType, &bid.AuthorId, &bid.Version, &bid.CreatedAt)
	if err != nil {
		return models.Bid{}, err
	}
	return bid, nil
}
