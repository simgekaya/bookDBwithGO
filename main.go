package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `gorm:"primaryKey"`
	ISBN   string `gorm:"unique"`
	Title  string `gorm:"not null"`
	Author string `gorm:"not null"`
}

func main() {
	// Gorm veri tabanı bağlantısı
	db, err := gorm.Open(sqlite.Open("Books.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// kitaplar tablosunu oluşturma
	err = db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal(err)
	}

	// veri girişi
	addBook := Book{ISBN: "98765412365", Title: "Masumiyet Müzesi", Author: "Orhan Pamuk"}
	err = addData(db, &addBook)
	if err != nil {
		log.Fatal(err)
	}
	//db.Create(&book1)
	//db.Create(&book2)

	// veri tabanı bağlantı kapatma
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}

func getData(db *gorm.DB) ([]Book, error) {
	var books []Book
	err := db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func addData(db *gorm.DB, book *Book) error {
	err := db.Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func updateData(db *gorm.DB, id uint, updatedData *Book) error {
	var book Book
	err := db.First(&book, id).Error
	if err != nil {
		return err
	}
	err = db.Model(&book).Updates(updatedData).Error
	if err != nil {
		return err
	}
	return nil
}

func deleteData(db *gorm.DB, id uint) error {
	var book Book
	err := db.First(&book, id).Error
	if err != nil {
		return err
	}
	err = db.Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
