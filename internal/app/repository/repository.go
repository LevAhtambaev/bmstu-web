package repository

import (
	"WAD-2022/internal/app/ds"
	"WAD-2022/internal/app/dsn"
	"context"
	"github.com/tjarratt/babble"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
)

type Repository struct {
	db *gorm.DB
}

func New(ctx context.Context) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetMangaByID(id uint) (*ds.Manga, error) {
	product := &ds.Manga{}

	err := r.db.First(product, "id = ?", "1").Error // find product with code D42
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *Repository) NewRandManga() (*ds.Manga, error) {
	babbler := babble.NewBabbler()
	babbler.Count = 1
	product := &ds.Manga{
		uint(rand.Intn(5000)),
		babbler.Babble(),
		uint(rand.Intn(2000)),
		babbler.Babble(),
		uint(rand.Intn(2000)),
	}
	err := r.db.Create(product)
	if err != nil {
		return nil, err.Error
	}
	return product, nil
}

func (r *Repository) CreateManga(product ds.Manga) error {
	return r.db.Create(product).Error
}
