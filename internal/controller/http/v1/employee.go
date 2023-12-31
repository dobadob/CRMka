package v1

import (
	"CRMka/internal/apperror"
	"CRMka/internal/domain/dto"
	"CRMka/internal/domain/entity"
	"CRMka/pkg/logging"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	employeesURL = "/employees"
	employeeURL  = "/employees/:id"
)

type Service interface {
	CreateEmployee(ctx context.Context, dto dto.CreateEmployeeDTO) (string, error)
	GetAllEmployees(ctx context.Context) (e []entity.Employee, err error)
	GetEmployeeByID(ctx context.Context, id string) (entity.Employee, error)
	UpdateEmployee(ctx context.Context, employee entity.Employee) error
	DeleteEmployeeByID(ctx context.Context, id string) error
}

type handler struct {
	logger  *logging.Logger
	service Service
}

func NewEmployeeHandler(l *logging.Logger, s Service) *handler {
	return &handler{
		logger:  l,
		service: s,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, employeesURL, apperror.Middleware(h.GetAllEmployees))
	router.HandlerFunc(http.MethodPost, employeesURL, apperror.Middleware(h.CreateEmployee))
	router.HandlerFunc(http.MethodGet, employeeURL, apperror.Middleware(h.GetEmployeeByUUID))
	router.HandlerFunc(http.MethodPut, employeeURL, apperror.Middleware(h.UpdateEmployee))
	router.HandlerFunc(http.MethodDelete, employeeURL, apperror.Middleware(h.DeleteEmployee))
}

func (h *handler) GetAllEmployees(w http.ResponseWriter, r *http.Request) error {
	employees, err := h.service.GetAllEmployees(context.Background())
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(employees)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

	return nil
}

func (h *handler) CreateEmployee(w http.ResponseWriter, r *http.Request) error {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	var newEmployee dto.CreateEmployeeDTO
	err = json.Unmarshal(bytes, newEmployee)
	if err != nil {
		return err
	}

	id, err := h.service.CreateEmployee(context.Background(), newEmployee)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
	return nil
}

func (h *handler) GetEmployeeByUUID(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is employee by id"))

	return nil
}

func (h *handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is update employee"))

	return nil
}

func (h *handler) PartiallyUpdateEmployee(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is partially update employee"))

	return nil
}

func (h *handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this delete employee"))

	return nil
}
