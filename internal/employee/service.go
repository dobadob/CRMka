package employee

import (
	"CRMka/pkg/logging"
	"context"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateEmployeeDTO) (e Employee, err error) {
	//TODO
	return
}
