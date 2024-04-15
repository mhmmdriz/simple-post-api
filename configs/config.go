package configs

import (
	"soal-eksplorasi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDBConnection() {
	dsn := "root:@tcp(127.0.0.1:3306)/post_category?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	initMigration()
}

func initMigration() {
	// DB.Migrator().DropTable(&models.Post{})
	// DB.Migrator().DropTable(&models.User{})
	// DB.Migrator().DropTable(&models.Category{})
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Category{})
}
