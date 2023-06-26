package models

import "gorm.io/gorm"

/*
Auto Migrate will create the tables if they don't exist
err = db.AutoMigrate(&Student{}, &Teacher{})              --> if multiple tables exist,
*/

func MigrateData(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
