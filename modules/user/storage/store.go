package userstorage

import "go.mongodb.org/mongo-driver/mongo"

type MgDBStorage struct {
	db *mongo.Client
}

func NewMgDBStorage(db *mongo.Client) *MgDBStorage {
	return &MgDBStorage{db: db}
}
