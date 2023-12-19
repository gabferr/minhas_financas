package config

import (
	"os"

	"github.com/gabferr/minhas_financas/schemas"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("Sqlite")
	dbPath := "./db/main.db"
	//Check if DB already exists
	_, error := os.Stat(dbPath)
	if os.IsNotExist(error) {
		logger.Info("database file not foun, crating ")
		//Create database and directory
		error := os.MkdirAll("./db", os.ModePerm)
		if error != nil {
			return nil, error
		}
		file, error := os.Create(dbPath)
		if error != nil {
			return nil, error
		}
		file.Close()
	}
	//Create db and Conect
	db, error := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if error != nil {
		logger.Errorf("Sqlite opening error: %v", error)
		return nil, error
	}
	//Migrate Schema
	error = db.AutoMigrate(&schemas.Opening{})
	if error != nil {
		logger.Errorf("Sqlite migration error: %v", error)
		return nil, error
	}
	//Return DB
	return db, nil
}
