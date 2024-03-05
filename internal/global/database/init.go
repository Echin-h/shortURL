package database

import (
	"ShortUrl/configs"
	"ShortUrl/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.Get().Mysql.Username,
		configs.Get().Mysql.Password,
		configs.Get().Mysql.Host,
		configs.Get().Mysql.Port,
		configs.Get().Mysql.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("linked database success!")
	panic(db.AutoMigrate(&model.ShortUrl{}))
}
