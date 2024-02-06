package datastore

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"myAPIProject/internal/apperrors"
	"myAPIProject/internal/config"
	"time"
)

type DB struct {
	Mongo      *mongo.Client
	Collection *mongo.Collection
}

func SetUpDatabase(cfg *config.Config, logger *logrus.Logger) (*DB, error) {
	if cfg.Mongo == nil {
		logger.Error("cfg.Mongo is nil. Cannot read config to init mongodb")
		return nil, &apperrors.NilMongoDBConfigErr
	}

	aa := uint64(100)
	bb := uint64(0)
	clientOptions := &options.ClientOptions{
		Hosts: []string{cfg.Mongo.DBHost},
		Auth: &options.Credential{
			Username: cfg.Mongo.DBUsername,
			Password: cfg.Mongo.DBPassword,
		},
		MinPoolSize: &bb,
		MaxPoolSize: &aa,
	}

	uri := cfg.Mongo.DBPrefix + "://" + cfg.Mongo.DBUsername + ":" + cfg.Mongo.DBPassword + "@" + cfg.Mongo.DBHost
	clientOptions.ApplyURI(uri)
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Infof("Connecting to MongoDB. Addr: %s://%s", cfg.Mongo.DBPrefix, cfg.Mongo.DBHost)
	client, err := mongo.Connect(ctxWithTimeout, clientOptions)
	if err != nil {
		logger.Error(err)
		return nil, apperrors.MongoDBInitErr.AppendMessage(err)
	}

	logger.Infof("Connecting to MongoDB. Addr: %s://%s", cfg.Mongo.DBPrefix, cfg.Mongo.DBHost)
	err = client.Ping(ctxWithTimeout, nil)
	if err != nil {
		logger.Error(err)
		return nil, apperrors.MongoDBInitErr.AppendMessage(err)
	}

	coll := client.Database(cfg.Mongo.DBName).Collection(cfg.Mongo.DBCollection)

	logger.Infof("MongoDB connection has been established. Host:%s://%s", cfg.Mongo.DBHost, cfg.Mongo.DBUsername)

	return &DB{
		Mongo:      client,
		Collection: coll}, nil
}
