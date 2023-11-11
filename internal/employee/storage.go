package employee

import "context"

type Storage interface {
	Create(ctx context.Context, employee Employee) (string, error)
	FindAll(ctx context.Context) (e []Employee, err error)
	FindOne(ctx context.Context, id string) (Employee, error)
	Update(ctx context.Context, employee Employee) error
	Delete(ctx context.Context, id string) error
}
