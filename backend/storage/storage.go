package storage

import "go.mongodb.org/mongo-driver/mongo"

type LocalInterface interface {
	Db() *mongo.Client
}
