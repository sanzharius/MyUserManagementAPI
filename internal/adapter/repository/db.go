package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"myAPIProject/internal/apperrors"
	"myAPIProject/internal/infrastructure/datastore"
	"myAPIProject/internal/usecase/repository"
	"time"
)

type dbRepository struct {
	collection *mongo.Collection
	db         *datastore.DB
	logger     *logrus.Logger
}

func NewDBRepository(collection *mongo.Collection, db *datastore.DB, logger *logrus.Logger) repository.DBRepository {
	return &dbRepository{collection: collection,
		db: db, logger: logger}
}

func (dr *dbRepository) Transaction(sessionFunc func(interface{}) (interface{}, error)) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	session, err := dr.db.Mongo.StartSession()
	if err != nil {
		dr.logger.Error(err)
		return nil, apperrors.MongoDBStartErr.AppendMessage(err)
	}

	defer session.EndSession(ctx)

	data, err := sessionFunc(session)
	return data, nil
}
