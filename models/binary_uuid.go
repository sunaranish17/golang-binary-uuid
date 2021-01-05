package models

import (
	"github.com/google/uuid"
)

//BinaryUUID -> binary uuid wrapper over uuid.UUID
type BinaryUUID uuid.UUID

func (b BinaryUUID) String() string {
	return uuid.UUID(b).String()
}

//MarshalJSON -> convert to json string
func (b BinaryUUID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(b)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}
