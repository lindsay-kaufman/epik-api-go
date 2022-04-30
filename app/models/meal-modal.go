package	models

import (
	"gorm.io/gorm"
	"github.com/lindsay-kaufman/go-epik-api/app/config"
	"time"
) 


type Meal struct {
	gorm.Model
	id int `gorm:"primaryKey"`
	name string
	notes string
	created_at time.Time
	user_id int
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Meal{})
}