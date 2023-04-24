package config

import (
	"example/go-mysql/models"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
*	DB Connection
 */
var Db *gorm.DB

/*
*	Initialize connection and migrate tables
 */
func InitDB() error {
	var err error
	sqlPass := viper.Get("SQL_PASS")
	sqlHost := viper.Get("SQL_HOST")
	sqlUser := viper.Get("SQL_USER")

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:3306)/?parseTime=True", sqlUser, sqlPass, sqlHost)

	Db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return err
	}

	Db.Exec("USE articles_db")

	// Migration to create tables for Order and Item schema
	Db.AutoMigrate(&models.Article{}, &models.Comment{})

	return nil
}
