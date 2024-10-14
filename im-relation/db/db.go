package db

import "gorm.io/gorm"

type DB struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) *DB {
	return &DB{db: db}
}

func (db *DB) Txn() *gorm.DB {
	return db.db.Begin()
}

func (db *DB) DB() *gorm.DB {
	return db.db
}
