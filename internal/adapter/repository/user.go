package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"myAPIProject/internal/apperrors"
	"myAPIProject/internal/domain/model"
	"myAPIProject/internal/infrastructure/datastore"
	"myAPIProject/internal/usecase/repository"
	"myAPIProject/internal/utils"
)

type userRepository struct {
	collection *mongo.Collection
	client     *datastore.DB
	logger     *logrus.Logger
}

func NewUserRepository(collection *mongo.Collection, client *datastore.DB, logger *logrus.Logger) repository.UserRepository {
	return &userRepository{collection: collection,
		client: client,
		logger: logger}
}

func (db *userRepository) Create(ctx context.Context, user *model.User) (primitive.ObjectID, error) {
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

func (db *userRepository) FindAll(ctx context.Context, paginationQuery *utils.PaginationQuery) ([]*model.User, error) {
	var res []*model.User
	findOptions := options.Find().SetSort(bson.D{{Key: paginationQuery.OrderBy, Value: 1}})
	findOptions.SetSkip(int64(paginationQuery.GetSkip()))
	findOptions.SetLimit(int64(paginationQuery.GetLimit()))
	cursor, err := db.collection.Find(ctx, bson.D{}, findOptions)
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
}

func (db *userRepository) FindUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	existingUser := &model.User{}
	err := db.collection.FindOne(ctx, bson.D{{"nickname", nickname}}).Decode(existingUser)
	if err == mongo.ErrNoDocuments {
		return nil, apperrors.MongoDBFindOneErr.AppendMessage(err)
	} else if err != nil {
		return nil, apperrors.MongoDBDataNotFoundErr.AppendMessage(err)
	}

	return existingUser, nil
}

func (db *userRepository) FindUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user := &model.User{}
	err := db.collection.FindOne(ctx, bson.D{{"_id,omitempty", userID}}).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, apperrors.MongoDBFindOneByIDErr.AppendMessage(err)
	} else if err != nil {
		return nil, apperrors.MongoDBDataNotFoundErr.AppendMessage(err)
	}

	return user, nil
}

func (db *userRepository) UpdateUser(ctx context.Context, user *model.User) (*mongo.UpdateResult, error) {
	filter := bson.D{{"_id,omitempty", user.ID}}
	userMarshalled, err := bson.Marshal(user)
	if err != nil {
		return nil, apperrors.UserBSONMarshalErr.AppendMessage(err)
	}

	update := bson.D{{"$set", bson.D{{"nickname", user.Nickname}, {"email", user.Email},
		{"first_name", user.FirstName}, {"last_name", user.LastName},
		{"password", user.Password}, {"created_at", user.Created.At},
		{"updated_at", user.UpdatedAt}, {"deleted_at", user.DeletedAt}}}}
	err = bson.Unmarshal(userMarshalled, update)

	result, err := db.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, apperrors.MongoDBUpdateErr.AppendMessage(err)
	}

	return result, nil
}

func (db *userRepository) DeleteOne(ctx context.Context, userID *uuid.UUID) error {
	filter := bson.D{{"_id", userID}}
	result, err := db.collection.DeleteOne(ctx, filter)
	if err != nil {
		return apperrors.MongoDBDeleteErr.AppendMessage(err)
	}
	if result.DeletedCount == 0 {
		return apperrors.MongoDBDeletedCountErr.AppendMessage(err)
	}

	return nil
}
