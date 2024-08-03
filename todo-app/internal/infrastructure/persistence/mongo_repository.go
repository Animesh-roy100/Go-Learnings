package persistence

import (
	"context"
	"time"
	"todo-app/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewMongoRepository(connectionString, database, collection string) (*MongoRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &MongoRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (r *MongoRepository) Save(todo *domain.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := r.client.Database(r.database).Collection(r.collection)
	_, err := coll.InsertOne(ctx, todo)
	return err
}

func (r *MongoRepository) FindAll() (*domain.TodoList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := r.client.Database(r.database).Collection(r.collection)
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	todoList := domain.NewTodoList()
	for cursor.Next(ctx) {
		var todo domain.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todoList.AddTodo(&todo)
	}

	return todoList, nil
}

func (r *MongoRepository) Update(todo *domain.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := r.client.Database(r.database).Collection(r.collection)
	filter := bson.M{"_id": todo.ID.String()}
	update := bson.M{"$set": bson.M{"completed": todo.Completed}}
	_, err := coll.UpdateOne(ctx, filter, update)
	return err
}
