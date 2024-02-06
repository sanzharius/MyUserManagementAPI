package registry

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"myAPIProject/internal/adapter/controller"
	"myAPIProject/internal/infrastructure/datastore"
)

type registry struct {
	collection *mongo.Collection
	db         *datastore.DB
	logger     *logrus.Logger
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *datastore.DB) Registry {
	return &registry{db: db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{UserController: r.NewUserController()}
}
