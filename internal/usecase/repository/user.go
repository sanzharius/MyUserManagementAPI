package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"myAPIProject/internal/domain/model"
	"myAPIProject/internal/utils"
)

/*type UserStorage struct {
	collection *mongo.Collection
	client     *mongo.Client
	logger     *logrus.Logger
}*/

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*uuid.UUID, error)
	FindAll(ctx context.Context, paginationQuery *utils.PaginationQuery) ([]*model.User, error)
	FindUserByNickname(ctx context.Context, nickname string) (*model.User, error)
	FindUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, userID *uuid.UUID) error
}

/*func (db *UserStorage) Create(ctx context.Context, user *model.User) (primitive.ObjectID, error) {
	result, err := db.collection.InsertOne(ctx, user)
	if err != nil {
		db.logger.Error(err)
		return primitive.NilObjectID, apperrors.MongoDBDataNotFoundErr.AppendMessage(err)
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		db.logger.Error(err)
		return primitive.NilObjectID, apperrors.MongoDBDataNotFoundErr.AppendMessage(err)
	}

	return id, nil
}

func (db *UserStorage) FindAll(ctx context.Context, filter bson.D) ([]*model.User, error) {
	var res []*model.User
	cursor, err := db.collection.Find(ctx, filter)
	if err != nil {
		db.logger.Error(err)
		return nil, apperrors.MongoDBFindErr.AppendMessage(err)
	}
	defer func() {
		_ = cursor.Close(ctx)
	}()

	err = cursor.All(ctx, &res)
	if err != nil {
		db.logger.Error(err)
		return nil, apperrors.MongoDBFindErr.AppendMessage(err)
	}

	if err := cursor.Err(); err != nil {
		db.logger.Error(err)
		return nil, apperrors.MongoDBFindErr.AppendMessage(err)
	}

	if len(res) <= 0 {
		return nil, apperrors.MongoDBDataNotFoundErr.AppendMessage(err)
	}

	return res, nil
}*/
