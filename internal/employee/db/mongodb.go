package db

import (
	"CRMka/internal/employee"
	"CRMka/pkg/logging"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, employee employee.Employee) (string, error) {
	d.logger.Debug("create employee")
	result, err := d.collection.InsertOne(ctx, employee)
	if err != nil {
		return "", fmt.Errorf("failed to create employee due to error: %v", err)
	}

	d.logger.Debug("convert IncertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(employee)
	return "", fmt.Errorf("dailed to convert objectid to hex. probably oid: %s", oid)
}

func (d *db) FindOne(ctx context.Context, id string) (e employee.Employee, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return e, fmt.Errorf("filed to convert hex to objectid. hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		//TODO 404
		return e, fmt.Errorf("failed to find one user by id: %s due to error: %v", id, err)
	}
	if err = result.Decode(&e); err != nil {
		return e, fmt.Errorf("failed to decode employee (id: %s) from DB due to error: %v", id, err)
	}

	return e, nil
}

func (d *db) Update(ctx context.Context, employee employee.Employee) error {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) employee.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
