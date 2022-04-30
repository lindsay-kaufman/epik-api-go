package config

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

 var (db * gorm.DB)

 func Connect(){
	 dbUrl := "host=localhost user=lindsaykaufman password=3139Lakeshore dbname=calendar port=5432"
	 d, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	 if err != nil {
		 panic(err)
	 }

	 db = d
 }

 func GetDB() gorm.DB{
	 return db
 }
