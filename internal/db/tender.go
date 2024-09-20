package db

import (
	"context"
	"tenderService/internal/models"
	"time"
)

func GetTenders(limit int, offset int, serviceTypes []string) ([]models.Tender, error) {
	var query string
	var args []interface{}

	// Базовый запрос
	query = `
		SELECT id, title, description, service_type, status, organization_id, version, created_at 
		FROM tenders`

	// Если есть фильтр по serviceTypes
	if len(serviceTypes) > 0 {
		query += `
		WHERE service_type = ANY($1)
		ORDER BY title ASC
		LIMIT $2 OFFSET $3;`
		args = append(args, serviceTypes, limit, offset)
	} else {
		query += `
		ORDER BY title ASC
		LIMIT $1 OFFSET $2;`
		args = append(args, limit, offset)
	}

	// Выполняем запрос
	rows, err := db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenders []models.Tender
	for rows.Next() {
		var tender models.Tender
		err := rows.Scan(&tender.Id, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Status, &tender.OrganizationId, &tender.Version, &tender.CreatedAt)
		if err != nil {
			return nil, err
		}
		tenders = append(tenders, tender)
	}
	return tenders, nil
}

func NewTender(title, desc, service_type, status, orgID, version string) (string, time.Time, error) {
	var tenderID string
	var createdAt time.Time
	err := db.QueryRow(context.Background(),
		`INSERT INTO tenders(
		title, description, service_type, status, organization_id, version)
		VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at;`,
		title, desc, service_type, status, orgID, version).Scan(&tenderID, &createdAt)

	if err != nil {
		return "", time.Now(), err
	}

	return tenderID, createdAt, nil
}

func AddTenderToUser(userID, tenderID string) error {
	_, err := db.Exec(context.Background(),
		`INSERT INTO user_tenders(user_id, tender_id)
		VALUES ($1, $2);`,
		userID, tenderID)
	return err
}

func GetUserTenders(userID string, limit, offset int) ([]models.Tender, error) {
	var tenders []models.Tender

	// Запрос
	query := `
		SELECT t.id, t.title, t.description, t.service_type, t.status, t.organization_id, t.version, t.created_at
		FROM tenders t
		JOIN user_tenders ut ON t.id = ut.tender_id
		WHERE ut.user_id = $1
		ORDER BY t.title ASC
		LIMIT $2 OFFSET $3
	`

	// Выполнение запроса
	rows, err := db.Query(context.Background(), query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Сканирование результатов
	for rows.Next() {
		var tender models.Tender
		err := rows.Scan(
			&tender.Id,
			&tender.Name,
			&tender.Description,
			&tender.ServiceType,
			&tender.Status,
			&tender.OrganizationId,
			&tender.Version,
			&tender.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		tenders = append(tenders, tender)
	}

	// Проверка на ошибки после выполнения запроса
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tenders, nil
}

func GetTenderStatus(tenderID string) (string, error) {
	var status string
	err := db.QueryRow(context.Background(),
		`SELECT status
		FROM tenders
		WHERE id = $1;`,
		tenderID).Scan(&status)
	if err != nil {
		return "", err
	}
	return status, nil
}

func EditTenderStatus(tenderID, status string) error {
	_, err := db.Exec(context.Background(),
		`UPDATE tenders
		SET status = $1
		WHERE id = $2;`,
		status, tenderID)
	return err
}

func GetTenderByID(tenderID string) (models.Tender, error) {
	var tender models.Tender

	// Запрос
	query := `
		SELECT t.id, t.title, t.description, t.service_type, t.status, t.organization_id, t.version, t.created_at
		FROM tenders t
		WHERE t.id = $1
	`	
	//TODO закончить запрос	
	// Выполнение запроса
	err := db.QueryRow(context.Background(), query, tenderID).Scan(
		&tender.Id,
		&tender.Name,
		&tender.Description,
		&tender.ServiceType,
		&tender.Status,
		&tender.OrganizationId,
		&tender.Version,
		&tender.CreatedAt,
	)
	
	return tender, err
}

