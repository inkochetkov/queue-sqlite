package sqlite

import (
	"database/sql"
	"log"
	p "path"

	"github.com/inkochetkov/exist"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB sqlite, orm gorm
type DB struct {
	db   *gorm.DB
	conn *sql.DB
}

// create connection database
func new(dataSourcePath string) *DB {

	dbGorm, err := gorm.Open(sqlite.Open(dataSourcePath), &gorm.Config{})
	if err != nil {
		log.Fatalf("SQL connect fail, %v", err)
	}

	conn, err := dbGorm.DB()
	if err != nil {
		log.Fatalf("SQL connect fail, %v", err)
	}

	err = dbGorm.AutoMigrate(&Entity{})
	if err != nil {
		log.Fatalf("SQL connect fail, %v", err)
	}

	return &DB{conn: conn, db: dbGorm}
}

// Start - get DataBase interface
func Start(path, fileName string) DataBase {

	dataSourcePath, err := checkFileBD(path, fileName)
	if err != nil {
		log.Fatalf("path to bd fail, %v", err)
		return nil
	}

	return new(dataSourcePath)
}

// checkFileBD - check dir and file exist
func checkFileBD(path, fileName string) (string, error) {

	url := p.Join(path, fileName)

	if ok := exist.CheckFile(url); ok {
		return url, nil
	}

	_, err := exist.InitDirFile(path, fileName)
	if err != nil {
		return "", err
	}

	return url, nil
}
