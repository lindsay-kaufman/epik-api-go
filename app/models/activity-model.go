package	models

import (
	"gorm.io/gorm"
	"github.com/lindsay-kaufman/go-epik-api/app/config"
	"time"
) 

var db *gorm.DB

type Activity struct {
	gorm.Model
	id int `gorm:"primaryKey"`
	description string
	created_at time.Time
	user_id int
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Activity{})
}