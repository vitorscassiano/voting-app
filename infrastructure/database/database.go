package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// type Database struct {
// 	conn *gorm.DB
// }

func InitializeRepositories() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("voting-app.db"), &gorm.Config{})
	if err != nil {
		panic("repositories connection does not start")
	}

	return db
}

// func (d *Database) Migrate(db *gorm.DB, entities ...interface{}) {
// 	d.conn.AutoMigrate(entities)
// }
