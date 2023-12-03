package mongo

import (
	"Task31a.3.1/pkg/storage"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// Store Хранилище данных в MongoDB.
type Store struct {
	client     mongo.Client
	collection *mongo.Collection
}

// New Конструктор объекта хранилища.
func New(connectionString string) (*Store, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("подключение к MongoDB провалено: %w", err)
	}
	collection := client.Database("testdb").Collection("posts")
	return &Store{
		client:     mongo.Client{},
		collection: collection,
	}, nil
}

// Posts вывод списка всех постов.
func (s *Store) Posts() ([]storage.Post, error) {
	cursor, err := s.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить поиск %w", err)
	}
	defer cursor.Close(context.Background())
	var results []storage.Post
	for cursor.Next(context.TODO()) {
		var elem storage.Post
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, fmt.Errorf("ошибка декодинга %w", err)
		}
		results = append(results, elem)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.TODO())
	return results, nil
}

// AddPost Добавление поста.
func (s *Store) AddPost(post storage.Post) error {
	post.CreatedAt = time.Now().Unix()
	post.PublishedAt = time.Now().Unix()
	_, err := s.collection.InsertOne(context.TODO(), post)
	if err != nil {
		return fmt.Errorf("ошибка добавления записи: %w", err)
	}
	return nil
}

// UpdatePost Обновление поста.
func (s *Store) UpdatePost(post storage.Post) error {
	filter := bson.M{"id": post.ID}
	update := bson.M{
		"$set": bson.M{
			"Title":       post.Title,
			"Content":     post.Content,
			"AuthorID":    post.AuthorID,
			"AuthorName":  post.AuthorName,
			"PublishedAt": time.Now().Unix(),
		},
	}
	_, err := s.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("ошибка обновления записи в MongoDB: %w", err)
	}
	return nil
}

// DeletePost Удаление поста.
func (s *Store) DeletePost(post storage.Post) error {
	filter := bson.D{{"id", post.ID}}
	_, err := s.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("ошибка при удалении записи в MongoDB: %w", err)
	}
	return nil
}
