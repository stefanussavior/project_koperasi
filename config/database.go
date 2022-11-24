package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(c *gin.Context) {
	dsn := "root@tcp(127.0.0.1:3306)/project_koperasi?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	c.JSON(200, gin.H{"Status": "Oke"})
}
