package models

type Books struct {
    // var-name for golang.			// var-name for postgres
	ID        uint    				`gorm:"primary key;autoIncrement" json:"id"`
	Author    *string 				`json:"author"`
	Title     *string 				`json:"title"`
	Publisher *string 				`json:"publisher"`
}
