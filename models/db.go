package models

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	DBPort, _ := viper.Get("DB.PORT").(string)
	DBHost, _ := viper.Get("DB.HOST").(string)
	DBUser, _ := viper.Get("DB.USERNAME").(string)
	DBPass, _ := viper.Get("DB.PASWORD").(string)
	DBName, _ := viper.Get("DB.NAME").(string)

	config := DBConfig{
		Host:     DBHost,
		Port:     DBPort,
		User:     DBUser,
		Password: DBPass,
		DB:       DBName,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DB)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
