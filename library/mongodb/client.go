package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/**
说明：
	连接 Mongodb 实例的连接池
*/
func ConnectDatabase(uri string, poolSize uint64) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(poolSize)) // 连接池

	if err != nil {
		return nil, err
	}
	return client, nil
}
