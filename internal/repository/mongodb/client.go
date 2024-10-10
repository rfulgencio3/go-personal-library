package mongodb

import (
    "context"
    "time"

    "github.com/rfulgencio3/go-personal-library/configs"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(config *configs.Config) (*mongo.Client, error) {
    clientOptions := options.Client().ApplyURI(config.MongoURI)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, err
    }
	
    if err := client.Ping(ctx, nil); err != nil {
        return nil, err
    }
    return client, nil
}
