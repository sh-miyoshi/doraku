package hobbydb

import (
	"context"
	"time"

	"github.com/sh-miyoshi/doraku/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBHandler is interface of dbHandler
type DBHandler interface {
	Initialize(mongoURL string) error
	GetHobbyByID(id int) (*HobbyDB, error)
}

type dbHandler struct {
	DBHandler
	client *mongo.Client
}

var inst = &dbHandler{}

// GetInst return instance of Database Handler
func GetInst() DBHandler {
	return inst
}

func (h *dbHandler) Initialize(mongoURL string) error {
	var err error
	// mongo.NewClient(url)
	h.client, err = mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = h.client.Connect(ctx)
	if err != nil {
		return err
	}

	err = h.client.Ping(ctx, nil)
    if err != nil {
        return err
    }


	logger.Info("Success to connect Mongo DB %s", mongoURL)
	return nil
}

func (h *dbHandler) GetHobbyByID(id int) (*HobbyDB, error) {
	collection := h.client.Database("doraku").Collection("hobby")
	res := collection.FindOne(context.Background(), bson.M{"id": id})
	var hobby HobbyDB
	if err := res.Decode(&hobby); err != nil {
		return nil, err
	}
	return &hobby, nil
}

// TODO: GetAllHobby
// TODO: GetHobbyByGroupNo
