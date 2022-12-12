package repository

import (
	"WAD-2022/internal/app/ds"
	"github.com/google/uuid"
	"log"
)

func (r *Repository) GetAllComics() ([]ds.Comics, error) {
	var comics []ds.Comics
	result := r.db.Order("uuid").Find(&comics)
	err := result.Error
	if err != nil {
		return comics, err
	}
	return comics, err
}

func (r *Repository) GetComicsByName(uuid string) (ds.Comics, error) {
	var comics ds.Comics
	result := r.db.First(&comics, "uuid = ?", uuid)
	err := result.Error
	return comics, err
}

func (r *Repository) GetComicsName(uuid string) (string, error) {
	var comics ds.Comics
	err := r.db.Select("name").First(&comics, "uuid = ?", uuid).Error
	return comics.Name, err
}

func (r *Repository) CreateComics(comics ds.Comics) error {
	err := r.db.Create(&comics).Error
	return err
}

func (r *Repository) ChangeDescription(uuid uuid.UUID, desc string) error {
	var comics ds.Comics
	comics.UUID = uuid
	result := r.db.Model(&comics).Update("Description", desc)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) DeleteComics(uuid string) (string, error) {
	var comics ds.Comics
	err := r.db.Where("uuid = ?", uuid).Delete(&comics).Error
	if err != nil {
		log.Print(err)
		return "no comics", err
	}
	return uuid, err
}

func (r *Repository) ChangeComics(uuid uuid.UUID, comics ds.Comics) (int, error) {
	comics.UUID = uuid
	err := r.db.Model(&comics).Updates(ds.Comics{Name: comics.Name,
		Year:        comics.Year,
		Price:       comics.Price,
		Rate:        comics.Rate,
		Genre:       comics.Genre,
		Volumes:     comics.Volumes,
		Description: comics.Description,
		Image:       comics.Image}).Error
	//if errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil {
		return 500, err
	}
	return 0, nil
}
