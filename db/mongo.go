package db

import (
	"context"
	"errors"

	mdb "go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo interface {
	Ping() error
	GetConnection() (*mdb.Database, error)
}

type mongo struct {
	uri    string
	ctx    context.Context
	dbName string
	client *mdb.Client
}

//NewMongo : instance of MongoDB
func NewMongo(ctx context.Context, uri string, dbName string) Mongo {
	return &mongo{ctx: ctx, uri: uri, dbName: dbName}
}

func (m *mongo) GetConnection() (*mdb.Database, error) {
	var err error
	if err = m.Ping(); err == nil {
		return m.client.Database(m.dbName), err
	}

	clientOpts := options.Client().ApplyURI(m.uri)
	m.client, err = mdb.Connect(m.ctx, clientOpts)

	if err != nil {
		return nil, err
	}

	return m.client.Database(m.dbName), err
}

func (m *mongo) Ping() error {
	if m.client == nil {
		return errors.New("Mongo Client Instance is nil")
	}
	return m.client.Ping(m.ctx, readpref.Primary())
}
