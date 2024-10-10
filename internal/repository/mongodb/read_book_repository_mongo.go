package mongodb

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rfulgencio3/go-personal-library/configs"
	"github.com/rfulgencio3/go-personal-library/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type readBookRepositoryMongo struct {
	collection *mongo.Collection
}

// NewReadBookRepository initializes the MongoDB repository for ReadBook.
func NewReadBookRepository(client *mongo.Client, config *configs.Config) *readBookRepositoryMongo {
	collection := client.Database(config.MongoDatabase).Collection(config.MongoCollection)
	return &readBookRepositoryMongo{collection: collection}
}

func (r *readBookRepositoryMongo) Create(readBook *domain.ReadBook) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	readBook.ID = uuid.New().String()

	_, err := r.collection.InsertOne(ctx, readBook)
	return err
}

func (r *readBookRepositoryMongo) GetByID(id string) (*domain.ReadBook, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var readBook domain.ReadBook
	filter := bson.M{"id": id}
	err := r.collection.FindOne(ctx, filter).Decode(&readBook)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}
	return &readBook, nil
}

// Similarmente, implemente os métodos Update, Delete e GetAll...
