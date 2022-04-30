package	models

import (
	"gorm.io/gorm"
	"github.com/lindsay-kaufman/go-epik-api/app/config"
	"time"
) 

type Mood struct {
	gorm.Model
	id int `gorm:"primaryKey"`
	notes string
	created_at time.Time
	user_id int
	score int
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Mood{})
}