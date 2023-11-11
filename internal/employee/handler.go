package employee

import (
	"CRMka/internal/apperror"
	"CRMka/internal/handlers"
	"CRMka/pkg/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	employeesURL = "/employees"
	employeeURL  = "/employees/:uuid"
)

type handler struct {
	name   string
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		name:   "employee",
		logger: logger,
	}
}

func (h *handler) GetName() string {
	return h.name
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, employeesURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, employeesURL, apperror.Middleware(h.CreateEmployee))
	router.HandlerFunc(http.MethodGet, employeeURL, apperror.Middleware(h.GetEmployeeByUUID))
	router.HandlerFunc(http.MethodPut, employeeURL, apperror.Middleware(h.UpdateEmployee))
	router.HandlerFunc(http.MethodPatch, employeeURL, apperror.Middleware(h.PartiallyUpdateEmployee))
	router.HandlerFunc(http.MethodDelete, employeeURL, apperror.Middleware(h.DeleteEmployee))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is list of employees"))

	return nil
}

func (h *handler) CreateEmployee(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("this is create employee"))

	return nil
}

func (h *handler) GetEmployeeByUUID(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is employee by uuid"))

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
