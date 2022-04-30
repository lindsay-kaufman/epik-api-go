package	models

import (
	"gorm.io/gorm"
	"github.com/lindsay-kaufman/go-epik-api/app/config"
	"time"
) 

type Goal struct {
	gorm.Model
	id int `gorm:"primaryKey"`
	title string
	description string
	completed bool
	start_date time.Time
	user_id int
	target_occurances int
	duration time.Duration
	created_at time.Time
	occurances int
	updated_at time.Time
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Goal{})
}