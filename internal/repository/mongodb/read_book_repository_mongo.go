package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rfulgencio3/go-personal-library/configs"
	"github.com/rfulgencio3/go-personal-library/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type readBookRepositoryMongo struct {
	collection *mongo.Collection
}

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

func (r *readBookRepositoryMongo) GetAll() ([]*domain.ReadBook, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var readBooks []*domain.ReadBook
	for cursor.Next(ctx) {
		var readBook domain.ReadBook
		if err := cursor.Decode(&readBook); err != nil {
			return nil, err
		}
		readBooks = append(readBooks, &readBook)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return readBooks, nil
}

func (r *readBookRepositoryMongo) Update(readBook *domain.ReadBook) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": readBook.ID}
	update := bson.M{
		"$set": bson.M{
			"book_id":           readBook.BookID,
			"start_date":        readBook.StartDate,
			"expected_end_date": readBook.ExpectedEndDate,
			"actual_end_date":   readBook.ActualEndDate,
			"rating":            readBook.Rating,
			"comments":          readBook.Comments, // Atualiza os comentários se necessário
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("read book not found")
	}

	return nil
}

func (r *readBookRepositoryMongo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("read book not found")
	}

	return nil
}

func (r *readBookRepositoryMongo) AddComment(id string, comment string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{
		"$push": bson.M{"comments": comment}, // Adiciona o comentário à lista existente
	}

	result, err := r.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(false))
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("read book not found")
	}

	return nil
}
