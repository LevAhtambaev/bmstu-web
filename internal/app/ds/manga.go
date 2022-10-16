package ds

import (
	"github.com/google/uuid"
)

type Manga struct {
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	Name        string
	Rate        uint64
	Year        uint64
	Genre       string
	Price       uint64
	Episodes    uint64
	Description string
}
