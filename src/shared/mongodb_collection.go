package shared

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBCollection interface {
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) MongoDBSingleResult
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

type mongoDBCollection struct {
	mongoCollection *mongo.Collection
}

type MongoDBSingleResult interface {
	Decode(v interface{}) error
}

func (m *mongoDBCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) MongoDBSingleResult {
	return m.mongoCollection.FindOne(ctx, filter, opts...)
}

func (m *mongoDBCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return m.mongoCollection.InsertOne(ctx, document, opts...)
}

func CreateMongoDBCollection() *mongoDBCollection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(fmt.Sprintf("could not connect to database: [%s]", err.Error()))
	}

	return &mongoDBCollection{
		mongoCollection: client.Database("marketingDB").Collection("mini-urls"),
	}
}
