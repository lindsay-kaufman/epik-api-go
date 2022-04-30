package	models

import (
	"gorm.io/gorm"
	"github.com/lindsay-kaufman/go-epik-api/app/config"
	"time"
) 


type ToDoList struct {
	gorm.Model
	id int `gorm:"primaryKey"`
	name string
	completed bool
	created_at time.Time
	user_id int
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&ToDoList{})
}