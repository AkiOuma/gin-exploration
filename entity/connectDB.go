package entity

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:000000@tcp(192.168.1.200:3306)/gundam?charset=utf8&parseTime=True"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: createLogger(),
	})
	if err != nil {
		panic(err)
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&GundamEntity{})
	DB.AutoMigrate(&PilotEntity{})
	DB.AutoMigrate(&OrganizationEntity{})
}

func createLogger() logger.Interface {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	return logger
}
