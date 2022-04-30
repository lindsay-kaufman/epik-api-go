package	models

import (
	"gorm.io/gorm"
	"github.com/lindsay-kaufman/go-epik-api/app/config"
	"time"
) 

type Sleep struct {
	gorm.Model
	id int `gorm:"primaryKey"`
	hours int
	quality string
	dreams string
	created_at time.Time
	user_id int
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Sleep{})
}