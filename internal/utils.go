package internal

import "github.com/google/uuid"

func GetNewUUID() (uuid.UUID, error) {
	var u uuid.UUID
	u, err := uuid.NewUUID()
	if err != nil {
		return u, err
	}
	return u, nil
}
