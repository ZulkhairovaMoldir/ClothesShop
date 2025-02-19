package config

import (
	"ClothesShop/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=hasan160 dbname=clothes_shop port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	log.Println("Подключение к базе данных установлено")

	// AutoMigrate to apply schema changes
	DB.AutoMigrate(&models.Product{})
}

// import (
//     "gorm.io/driver/postgres"
//     "gorm.io/gorm"
//     "log"
//     "os"
// )

// var DB *gorm.DB

// func InitDB() {
//     // Use environment variables or hardcoded DSN for now
//     dsn := os.Getenv("jdbc:postgresql://localhost:5432/clothes_shop")
//     if dsn == "" {
//         dsn = "host=localhost user=manager password=manager dbname=clothes_shop port=5432 sslmode=disable"
//     }

//     var err error
//     DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
//     if err != nil {
//         log.Fatalf("Failed to connect to the database: %v", err)
//     }
//     log.Println("Connected to the database successfully")
// }
