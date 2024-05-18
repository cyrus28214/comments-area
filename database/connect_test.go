package database

import "testing"

// package database

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type DatabaseConfig struct {
// 	Host     string `mapstructure:"host"`
// 	Port     string `mapstructure:"port"`
// 	User     string `mapstructure:"user"`
// 	Password string `mapstructure:"password"`
// 	Database string `mapstructure:"database"`
// }

// func Connect(config DatabaseConfig) (*gorm.DB, error) {
// 	dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

func TestConnect(t *testing.T) {
	db, err := Connect(DatabaseConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "test",
		Password: "test",
		Database: "test",
	})
	if err != nil {
		t.Error(err)
	}
	sql, err := db.DB()
	if err != nil {
		t.Error(err)
	}
	err = sql.Ping()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Connect success")
}
