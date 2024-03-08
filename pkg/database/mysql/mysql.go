package mysql

import (
	"CodegreeWebbs/pkg/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.LoadDataSourceName()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal(err)

	}

	return db

}
