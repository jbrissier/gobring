package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Bring struct {
	gorm.Model
	ID          uint   `json:"id"`
	Where       string `json:"where"`
	Until       uint   `json:"until"`
	User        string `json:"user"`
	Description string `json:"description"`
}

type BringItem struct {
	gorm.Model
	Bring       *Bring `gorm:"foreignKey:BringRef"`
	BringRef    int
	User        string
	Description string
}

var db *gorm.DB

func GetDB() (*gorm.DB, error) {

	// make the db connection singelton ? is this a good pattern

	if db != nil {
		return db, nil
	}

	newDb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db = newDb

	db.AutoMigrate(&Bring{})
	db.AutoMigrate(&BringItem{})

	return db, nil

}

func (b *Bring) Save() error {

	db, err := GetDB()
	if err != nil {
		return err
	}

	db.Save(b)
	return nil
}

func FindBring(id uint) (*Bring, error) {

	var bring Bring

	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	res := db.Find(&Bring{ID: id}).First(&bring)
	if res.Error != nil {
		return nil, res.Error
	}

	return &bring, nil

}
func DeleteBring(int uint) error {

	db, err := GetDB()
	if err != nil {
		return err
	}
	db.Delete(&Bring{ID: int})

	return nil
}

func GetAllBrings() ([]Bring, error) {

	var brings []Bring

	db, err := GetDB()

	if err != nil {
		return nil, err
	}
	db.Find(&brings)

	return brings, nil
}
