package project

import (
	"CRMka/internal/handlers"
	"CRMka/pkg/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	projectsURL = "/projects"
	projectURL  = "/projects/:uuid"
)

type handler struct {
	name   string
	logger logging.Logger
}

func NewHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		name:   "project",
		logger: logger,
	}
}

func (h *handler) GetName() string {
	return h.name
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(projectsURL, h.GetList)
	router.POST(projectsURL, h.CreateProject)
	router.GET(projectURL, h.GetProjectByUUID)
	router.PUT(projectURL, h.UpdateProject)
	router.PATCH(projectURL, h.PartiallyUpdateProject)
	router.DELETE(projectURL, h.DeleteProject)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is list of projects"))
}

func (h *handler) CreateProject(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is create project"))
}

func (h *handler) GetProjectByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is project by uuid"))
}

func (h *handler) UpdateProject(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is update project"))
}

func (h *handler) PartiallyUpdateProject(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is partially update project"))
}

func (h *handler) DeleteProject(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this delete project"))
}
