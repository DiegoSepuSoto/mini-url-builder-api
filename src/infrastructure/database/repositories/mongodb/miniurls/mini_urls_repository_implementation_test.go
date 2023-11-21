package miniurls

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

type mongoCollectionMock struct {
	mock.Mock
}

func (m *mongoCollectionMock) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document, opts)

	if args.Get(1) == nil {
		return args.Get(0).(*mongo.InsertOneResult), nil
	}

	return nil, args.Error(1)
}

func (m *mongoCollectionMock) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) shared.MongoDBSingleResult {
	args := m.Called(ctx, filter, opts)

	return args.Get(0).(shared.MongoDBSingleResult)
}

type mongoSingleResultMock struct {
	mock.Mock
}

func (sr *mongoSingleResultMock) Decode(v interface{}) error {
	args := sr.Called(v)

	if args.Get(0) == nil {
		return nil
	}

	return args.Error(0)
}
