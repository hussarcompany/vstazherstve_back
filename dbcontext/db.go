package dbcontext

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Структура представляющая подключение к базе данных
type DB struct {
	db *mongo.Database
}

// Возвращает новый экземпляр подключения к базе данных
func New(db *mongo.Database) *DB {
	return &DB{db}
}

func (db *DB) DB() *mongo.Database {
	return db.db
}
