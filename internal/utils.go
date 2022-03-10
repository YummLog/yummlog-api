package internal

import "github.com/google/uuid"

func GetNewUUID() (string, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
