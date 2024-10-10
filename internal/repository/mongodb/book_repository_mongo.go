package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rfulgencio3/go-personal-library/configs" // Atualizado para importar corretamente
	"github.com/rfulgencio3/go-personal-library/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// bookRepositoryMongo is the struct that implements the repository.BookRepository interface for MongoDB.
type bookRepositoryMongo struct {
	collection *mongo.Collection
}

// NewBookRepository creates a new book repository using MongoDB.
func NewBookRepository(client *mongo.Client, config *configs.Config) *bookRepositoryMongo {
	collection := client.Database(config.MongoDatabase).Collection(config.MongoCollection)
	return &bookRepositoryMongo{
		collection: collection,
	}
}

// Create inserts a new book into the MongoDB collection.
func (r *bookRepositoryMongo) Create(book *domain.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Generate a new UUID for the book ID
	book.ID = uuid.New().String()

	_, err := r.collection.InsertOne(ctx, book)
	return err
}

// GetByID retrieves a book by its ID from the MongoDB collection.
func (r *bookRepositoryMongo) GetByID(id string) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid book ID")
	}

	var book domain.Book
	filter := bson.M{"_id": objID}
	err = r.collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return &book, nil
}

// Update modifies an existing book in the MongoDB collection.
func (r *bookRepositoryMongo) Update(book *domain.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(book.ID)
	if err != nil {
		return errors.New("invalid book ID")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"title":     book.Title,
			"subtitle":  book.Subtitle,
			"author":    book.Author,
			"pages":     book.Pages,
			"publisher": book.Publisher,
			"comments":  book.Comments,
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("book not found")
	}

	return nil
}

// Delete removes a book from the MongoDB collection by its ID.
func (r *bookRepositoryMongo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid book ID")
	}

	filter := bson.M{"_id": objID}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("book not found")
	}

	return nil
}

// GetAll retrieves all books from the MongoDB collection.
func (r *bookRepositoryMongo) GetAll() ([]*domain.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var books []*domain.Book
	for cursor.Next(ctx) {
		var book domain.Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
