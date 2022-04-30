package	models

import (
	"gorm.io/gorm"
	"github.com/lindsay-kaufman/go-epik-api/app/config"
	"time"
) 

//do i not need to declare db here since i already did in the activity file?
//var db *gorm.DB

type User struct {
	gorm.Model
	firstname string 
	lastname string 
	email string 
	password string 
	created_at time.Time 
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}