package database

import "testing"

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
