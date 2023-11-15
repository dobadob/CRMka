package composites

import (
	"CRMka/internal/adapters/db/mongodb"
	"CRMka/internal/controller/http/v1"
	"CRMka/internal/domain/service"
	"CRMka/pkg/logging"
)

type EmployeeComposite struct {
	Storage service.Storage
	Service v1.Service
	Handler v1.Handler
}

func NewEmployeeComposite(logger *logging.Logger, mongoComposite *MongoDBComposite) *EmployeeComposite {
	storage := mongodb.NewEmployeeStorage(logger, mongoComposite.db)
	service := service.NewEmployeeService(logger, storage)
	handler := v1.NewEmployeeHandler(logger, service)
	return &EmployeeComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}
}
