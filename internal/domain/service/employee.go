package service

import (
	"CRMka/internal/controller/http/dto"
	"CRMka/internal/domain/entity"
	"CRMka/pkg/logging"
	"context"
)

type Storage interface {
	Create(ctx context.Context, e dto.CreateEmployeeDTO) (string, error)
	GetAll(ctx context.Context) (e []entity.Employee, err error)
	GetOne(ctx context.Context, id string) (e entity.Employee, err error)
	Update(ctx context.Context, e entity.Employee) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	storage Storage
	logger  *logging.Logger
}

func NewService(logger *logging.Logger, storage Storage) *service {
	return &service{
		storage: storage,
		logger:  logger,
	}
}

func (s service) CreateEmployee(ctx context.Context, dto dto.CreateEmployeeDTO) (string, error) {
	return s.storage.Create(ctx, dto)
}

func (s service) GetAllEmployees(ctx context.Context) (e []entity.Employee, err error) {
	return s.storage.GetAll(ctx)
}

func (s service) GetEmployeeByID(ctx context.Context, id string) (entity.Employee, error) {
	return s.storage.GetOne(ctx, id)
}

func (s service) UpdateEmployee(ctx context.Context, e entity.Employee) error {
	return s.storage.Update(ctx, e)
}

func (s service) DeleteEmployeeByUUID(ctx context.Context, id string) error {
	return s.storage.Delete(ctx, id)
}
