package employee

import (
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
	logger logging.Logger
}

func NewHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		name:   "employee",
		logger: logger,
	}
}

func (h *handler) GetName() string {
	return h.name
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(employeesURL, h.GetList)
	router.POST(employeesURL, h.CreateEmployee)
	router.GET(employeeURL, h.GetEmployeeByUUID)
	router.PUT(employeeURL, h.UpdateEmployee)
	router.PATCH(employeeURL, h.PartiallyUpdateEmployee)
	router.DELETE(employeeURL, h.DeleteEmployee)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is list of employees"))
}

func (h *handler) CreateEmployee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is create employee"))
}

func (h *handler) GetEmployeeByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is employee by uuid"))
}

func (h *handler) UpdateEmployee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is update employee"))
}

func (h *handler) PartiallyUpdateEmployee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is partially update employee"))
}

func (h *handler) DeleteEmployee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this delete employee"))
}
