package model

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Bring struct {
	gorm.Model
	ID          uint   `json:"id",gorm:"primaryKey"`
	Where       string `json:"where"`
	Until       uint   `json:"until"`
	User        string `json:"user"`
	Description string `json:"description"`
}

type BringItem struct {
	gorm.Model
	ID          uint   `json:"id",gorm:"primaryKey`
	Bring       *Bring `gorm:"foreignKey:BringRef"`
	BringRef    int
	User        string `json:"user"`
	Description string `json:"description"`
}

type BringDB struct {
	db         *gorm.DB
	DBLocation string
}

func NewBringDB(location string) *BringDB {
	return &BringDB{
		DBLocation: location,
	}
}

func (b *BringDB) GetDB() (*gorm.DB, error) {

	// make the db connection singelton ? is this a good pattern

	if b.db != nil {
		return b.db, nil
	}

	newDb, err := gorm.Open(sqlite.Open(b.DBLocation), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	b.db = newDb

	b.db.AutoMigrate(&Bring{})
	b.db.AutoMigrate(&BringItem{})

	return b.db, nil

}

func (b *BringDB) SaveBring(bring *Bring) error {

	db, err := b.GetDB()
	if err != nil {
		return err

	}

	db.Save(bring)
	return nil

}

func (b *BringDB) FindBring(id uint) (*Bring, error) {

	var bring Bring

	db, err := b.GetDB()
	if err != nil {
		return nil, err
	}
	res := db.Find(&Bring{ID: id}).First(&bring)
	if res.Error != nil {
		return nil, res.Error
	}

	return &bring, nil

}
func (b *BringDB) DeleteBring(int uint) error {

	db, err := b.GetDB()
	if err != nil {
		return err
	}
	db.Delete(&Bring{ID: int})

	return nil
}

func (b *BringDB) GetAllBrings() ([]Bring, error) {

	var brings []Bring

	db, err := b.GetDB()

	if err != nil {
		return nil, err
	}
	db.Find(&brings)

	return brings, nil
}

func (b *BringDB) GetItems(bring *Bring) ([]BringItem, error) {

	db, err := b.GetDB()

	var items = []BringItem{}

	if err != nil {
		return nil, err
	}
	db.Find(&items, "bring_ref = ?", bring.ID)
	return items, nil
}

func (br *BringDB) AddBrintItem(b *Bring, item *BringItem) error {

	db, err := br.GetDB()

	if err != nil {
		return err
	}
	item.Bring = b
	db.Create(item)

	return nil
}

func (br *BringDB) DeleteBringItem(id uint) error {

	db, err := br.GetDB()

	if err != nil {
		return err
	}
	fmt.Printf("delete item with id %d", id)
	res := db.Delete(&BringItem{}, id)
	if res.Error != nil {
		fmt.Printf("Error: %s", res.Error)
		return res.Error
	}
	return nil
}
