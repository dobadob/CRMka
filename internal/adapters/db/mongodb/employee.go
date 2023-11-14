package mongodb

import (
	"CRMka/internal/apperror"
	"CRMka/internal/controller/http/dto"
	"CRMka/internal/domain/entity"
	"CRMka/pkg/logging"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type employeeStorage struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) *employeeStorage {
	return &employeeStorage{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

func (d *employeeStorage) Create(ctx context.Context, e dto.CreateEmployeeDTO) (string, error) {
	d.logger.Debug("create employee")
	result, err := d.collection.InsertOne(ctx, e)
	if err != nil {
		return "", fmt.Errorf("failed to create employee due to error: %v", err)
	}

	d.logger.Debug("convert IncertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(e)
	return "", fmt.Errorf("dailed to convert objectid to hex. probably oid: %s", oid)
}

func (d *employeeStorage) GetAll(ctx context.Context) (e []entity.Employee, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return e, fmt.Errorf("failed to find all employyes due to error: %v", err)
	}

	if err := cursor.All(ctx, &e); err != nil {
		return e, fmt.Errorf("failed to read all documents from cursor. error: %v", err)
	}

	return e, nil
}

func (d *employeeStorage) GetOne(ctx context.Context, id string) (e entity.Employee, err error) {
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

func (d *employeeStorage) Update(ctx context.Context, e entity.Employee) error {
	oid, err := primitive.ObjectIDFromHex(e.ID)
	if err != nil {
		return fmt.Errorf("failed to convert employee ID to ObjectID. ID=%s", e.ID)
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

func (d *employeeStorage) Delete(ctx context.Context, id string) error {
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
