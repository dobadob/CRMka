package company

import (
	"CRMka/internal/handlers"
	"CRMka/pkg/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	companiesURL = "/companies"
	companyURL   = "/companies/:uuid"
)

type handler struct {
	name   string
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		name:   "company",
		logger: logger,
	}
}

func (h *handler) GetName() string {
	return h.name
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(companiesURL, h.GetList)
	router.POST(companiesURL, h.CreateCompany)
	router.GET(companyURL, h.GetCompanyByUUID)
	router.PUT(companyURL, h.UpdateCompany)
	router.PATCH(companyURL, h.PartiallyUpdateCompany)
	router.DELETE(companyURL, h.DeleteCompany)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is list of companies"))
}

func (h *handler) CreateCompany(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is create company"))
}

func (h *handler) GetCompanyByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is company by uuid"))
}

func (h *handler) UpdateCompany(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is update company"))
}

func (h *handler) PartiallyUpdateCompany(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is partially update company"))
}

func (h *handler) DeleteCompany(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this delete company"))
}
