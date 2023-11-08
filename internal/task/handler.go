package task

import (
	"CRMka/internal/handlers"
	"CRMka/pkg/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	tasksURL = "/tasks"
	taskURL  = "/tasks/:uuid"
)

type handler struct {
	name   string
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		name:   "task",
		logger: logger,
	}
}

func (h *handler) GetName() string {
	return h.name
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(tasksURL, h.GetList)
	router.POST(tasksURL, h.CreateTask)
	router.GET(taskURL, h.GetTaskByUUID)
	router.PUT(taskURL, h.UpdateTask)
	router.PATCH(taskURL, h.PartiallyUpdateTask)
	router.DELETE(taskURL, h.DeleteTask)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is list of tasks"))
}

func (h *handler) CreateTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is create task"))
}

func (h *handler) GetTaskByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is task by uuid"))
}

func (h *handler) UpdateTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is update task"))
}

func (h *handler) PartiallyUpdateTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is partially update task"))
}

func (h *handler) DeleteTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this delete task"))
}
