package databases

import (
	"bubble/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() (err error) {
	dsn := "root:Hjy20030920@@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	//模型绑定
	err = DB.AutoMigrate(&models.Todo{})
	if err != nil {
		return err
	}
	return nil
}
