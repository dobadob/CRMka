package db

import (
	"CRMka/internal/apperror"
	"CRMka/internal/employee"
	"CRMka/pkg/logging"
	"context"
	"errors"
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

func (d *db) FindAll(ctx context.Context) (e []employee.Employee, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return e, fmt.Errorf("failed to find all employyes due to error: %v", err)
	}

	if err := cursor.All(ctx, &e); err != nil {
		return e, fmt.Errorf("failed to read all documents from cursor. error: %v", err)
	}

	return e, nil
}

func (d *db) FindOne(ctx context.Context, id string) (e employee.Employee, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return e, fmt.Errorf("filed to convert hex to objectid. hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return e, apperror.ErrNotFound
		}
		return e, fmt.Errorf("failed to find one user by id: %s due to error: %v", id, err)
	}
	if err = result.Decode(&e); err != nil {
		return e, fmt.Errorf("failed to decode employee (id: %s) from DB due to error: %v", id, err)
	}

	return e, nil
}

func (d *db) Update(ctx context.Context, e employee.Employee) error {
	oid, err := primitive.ObjectIDFromHex(e.Id)
	if err != nil {
		return fmt.Errorf("failed to convert employee ID to ObjectID. ID=%s", e.Id)
	}

	filter := bson.M{"_id": oid}

	employeeBytes, err := bson.Marshal(e)
	if err != nil {
		return fmt.Errorf("failed to marshal employee. error: %v", err)
	}

	var updateEmployeeObj bson.M
	err = bson.Unmarshal(employeeBytes, updateEmployeeObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal employee bytes. error: %v", err)
	}

	delete(updateEmployeeObj, "_id")

	update := bson.M{"$set": updateEmployeeObj}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute update employee query. error: %v", err)
	}

	if result.MatchedCount == 0 {
		return apperror.ErrNotFound
	}

	d.logger.Tracef("Matched %d document and Modified %d document", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert employee ID to ObjectID. ID=%s", id)
	}

	filter := bson.M{"_id": oid}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %v", err)
	}
	if result.DeletedCount == 0 {
		return apperror.ErrNotFound
	}

	d.logger.Tracef("Deleted %d document", result.DeletedCount)

	return nil
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) employee.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
