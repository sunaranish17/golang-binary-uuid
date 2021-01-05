package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID      BinaryUUID `gorm:"primary_key"`
	Name    string     `gorm:"not null"`
	Comment Comment
}

//Before Create ->  runs before creating the model
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	p.ID = BinaryUUID(id)
	return err
}

//Comment is the model for comment table
type Comment struct {
	ID     uuid.UUID `gorm:"type:binary(16);prinary_key;default:(UUID_TO_BIN(UUID()));"`
	Name   string    `gorm:"size:128;not null;"`
	PostID BinaryUUID
}
