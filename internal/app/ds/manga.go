package ds

type Manga struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Year  uint
	Genre string
	Price uint
}
