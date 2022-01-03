package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	BuildingName string `gorm:"size:255"`
	RentFee      int    `gorm:"size:255"`
	Address      string `gorm:"size:255"`
}

type Db struct {
	DB *gorm.DB
}

func (d *Db) DbInit() {
	var err error
	d.DB, err = gorm.Open(mysql.Open("gorm:password@/test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("Error during initializing database")
	}
	log.Println("Database is initialized !")
}

func (d *Db) DbMigration() {
	d.DB.AutoMigrate(&Property{})
	log.Println("Auto migration end !")
}

func (d *Db) DbInsert(building_name string, rent_fee int, address string) {
	d.DB.Create(&Property{
		BuildingName: building_name,
		RentFee: rent_fee,
		Address: address,
	})
}
