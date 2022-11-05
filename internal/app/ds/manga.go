package ds

import (
	"github.com/google/uuid"
)

type Manga struct {
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	Name        string
	Rate        float64
	Year        uint64
	Genre       string
	Price       uint64
	Volumes     uint64
	Description string
	Image       string
}
