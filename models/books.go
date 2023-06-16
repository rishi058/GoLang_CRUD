package models

import "gorm.io/gorm"

type BooksTable struct {
    // var-name for golang.			// var-name for postgres
	ID        uint    				`gorm:"primary key;autoIncrement" json:"id"`
	Author    *string 				`json:"author"`
	Title     *string 				`json:"title"`
	Publisher *string 				`json:"publisher"`
}

/*
Auto Migrate will create the tables if they don't exist
err = db.AutoMigrate(&Student{}, &Teacher{})              --> if multiple tables exist,
*/

func MigrateData(db *gorm.DB) error {
	err := db.AutoMigrate(&BooksTable{})
	return err
}
