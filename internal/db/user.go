package db

import (
	"context"
	"fmt"
)

// GetIDByUsername возвращает идентификатор пользователя по его имени пользователя.
//
// Параметр:
//  - username: имя пользователя.
//
// Возвращает:
//  - идентификатор пользователя в виде строки.
//  - ошибка, если пользователь не найден.
func GetIDByUsername(username string) (string, error) {
	var userID string
    err := db.QueryRow(context.Background(),
	`SELECT id FROM employee WHERE username = $1`, username).Scan(&userID)
	if err != nil {
		return "", fmt.Errorf("failed to find user by username: %v", err)
	}
	return userID, nil
}
// IsResponsible checks if a user is responsible for an organization.
//
// Parameters:
// - userID: the ID of the user.
// - orgID: the ID of the organization.
//
// Returns:
// - bool: true if the user is responsible for the organization, false otherwise.
// - error: an error if the check fails.
func IsResponsible(userID string, orgID string) (bool, error) {
	var isResponsible bool
	err := db.QueryRow(context.Background(),
		`SELECT EXISTS(SELECT 1 FROM organization_responsible WHERE organization_id = $1 AND user_id = $2)`,
		orgID, userID).Scan(&isResponsible)
	if err != nil {
		return false, fmt.Errorf("failed to check user responsibility: %v", err)
	}
	return isResponsible, nil
}
