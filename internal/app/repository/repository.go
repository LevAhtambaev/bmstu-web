package repository

import (
	"WAD-2022/internal/app/ds"
	"WAD-2022/internal/app/dsn"
	"context"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func (r *Repository) GetAllManga() ([]ds.Manga, error) {
	var manga []ds.Manga
	result := r.db.Find(&manga)
	err := result.Error
	if err != nil {
		return manga, err
	}
	return manga, err
}

func (r *Repository) GetMangaByName(name string) ([]ds.Manga, error) {
	var manga []ds.Manga
	result := r.db.Where("name = ?", name).Find(&manga)
	err := result.Error
	if err != nil {
		return manga, err
	}
	return manga, err
}

func (r *Repository) CreateManga(manga ds.Manga) error {
	err := r.db.Create(&manga).Error
	return err
}

func (r *Repository) ChangeDescription(uuid uuid.UUID, desc string) error {
	var manga ds.Manga
	manga.UUID = uuid
	result := r.db.Model(&manga).Update("Description", desc)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) DeleteManga(uuid string) error {
	var manga ds.Manga
	result := r.db.Delete(&manga, "uuid = ?", uuid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
