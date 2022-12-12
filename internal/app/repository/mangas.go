package repository

import (
	"WAD-2022/internal/app/ds"
	"github.com/google/uuid"
	"log"
)

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

func (r *Repository) GetMangaName(uuid string) (string, error) {
	var manga ds.Manga
	err := r.db.Select("name").First(&manga, "uuid = ?", uuid).Error
	return manga.Name, err
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
	//res := r.db.First(&manga, uuid)
	//if res.Error != nil {
	//	return "no such rows", res.Error
	//}
	result := r.db.Delete(&manga, "uuid = ?", uuid)
	log.Print(result.Error)
	if result.Error != nil {
		return "no manga", result.Error
	}
	return uuid, result.Error
}

func (r *Repository) ChangeManga(uuid uuid.UUID, manga ds.Manga) (int, error) {
	manga.UUID = uuid
	err := r.db.Model(&manga).Updates(ds.Manga{Name: manga.Name,
		Year:        manga.Year,
		Price:       manga.Price,
		Rate:        manga.Rate,
		Genre:       manga.Genre,
		Volumes:     manga.Volumes,
		Description: manga.Description,
		Image:       manga.Image}).Error
	//if errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil {
		return 500, err
	}
	return 0, nil
}
