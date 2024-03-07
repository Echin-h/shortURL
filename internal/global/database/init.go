package database

import (
	"ShortUrl/configs"
	"ShortUrl/internal/global/log"
	"ShortUrl/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type URL struct {
	*gorm.DB
}

var Url = &URL{}

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.Get().Mysql.Username,
		configs.Get().Mysql.Password,
		configs.Get().Mysql.Host,
		configs.Get().Mysql.Port,
		configs.Get().Mysql.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.SugarLogger.Error("open database error: ", err)
		return
	}
	fmt.Println("linked database success!")
	err1 := db.AutoMigrate(&model.ShortUrl{})
	if err1 != nil {
		log.SugarLogger.Error("auto migrate error: ", err1)
		return
	}
	Url = &URL{db}
}
