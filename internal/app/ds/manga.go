package ds

import (
	"github.com/google/uuid"
)

type Manga struct {
	UUID        uuid.UUID `db:"uuid" gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	Name        string    `db:"name"`
	Rate        float64   `db:"rate"`
	Year        uint64    `db:"year"`
	Genre       string    `db:"genre"`
	Price       uint64    `db:"price"`
	Volumes     uint64    `db:"volumes"`
	Description string    `db:"description"`
	Image       string    `db:"image"`
}

func (Manga) TableName() string {
	return "manga"
}
