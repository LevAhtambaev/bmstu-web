package repository

import (
	"WAD-2022/internal/app/ds"
	"WAD-2022/internal/app/dsn"
	"context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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

func (r *Repository) GetMangaByName(uuid string) (ds.Manga, error) {
	var manga ds.Manga
	result := r.db.First(&manga, "uuid = ?", uuid)
	err := result.Error
	return manga, err
}

func (r *Repository) GetCar(uuid uuid.UUID) (ds.Manga, error) {
	var car ds.Manga
	err := r.db.First(&car, uuid).Error
	return car, err
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

func (r *Repository) DeleteManga(uuid string) (string, error) {
	var manga ds.Manga
	res := r.db.First(&manga, "uuid = ?", uuid)
	if res.Error != nil {
		return "no such rows", res.Error
	}
	result := r.db.Delete(&manga, "uuid = ?", uuid)
	log.Print(result.Error)
	if result.Error != nil {
		return "no manga", result.Error
	}
	return uuid, result.Error
}

func (r *Repository) GetCart() ([]ds.Cart, error) {
	var cart []ds.Cart
	err := r.db.Find(&cart).Error
	return cart, err
}

func (r *Repository) AddToCart(cart ds.Cart) error {
	err := r.db.Create(&cart).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteFromCart(manga uuid.UUID) (int, error) {
	var cart ds.Cart
	err := r.db.Where("manga = ?", manga).Delete(&cart).Error
	if err != nil {
		return 500, err
	}
	return 0, nil
}
