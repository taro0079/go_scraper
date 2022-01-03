package database

import "github.com/jinzhu/gorm"

type Property struct {
	gorm.Model
	BuildingName string  `gorm:"size:255"`
	RentFee      float64 `gorm:"size:255"`
}

func (p Property) DbInit() {
	db, err := gorm.Open("mysql", "root@/sample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error during initializing database")
	}
	defer db.Close()
}
