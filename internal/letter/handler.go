package letter

import (
	"CRMka/internal/handlers"
	"CRMka/pkg/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	lettersURL = "/letters"
	letterURL  = "/letters/:uuid"
)

type handler struct {
	name   string
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		name:   "letter",
		logger: logger,
	}
}

func (h *handler) GetName() string {
	return h.name
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(lettersURL, h.GetList)
	router.POST(lettersURL, h.CreateLetter)
	router.GET(letterURL, h.GetLetterByUUID)
	router.PUT(letterURL, h.UpdateLetter)
	router.PATCH(letterURL, h.PartiallyUpdateLetter)
	router.DELETE(letterURL, h.DeleteLetter)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is list of letters"))
}

func (h *handler) CreateLetter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is create letter"))
}

func (h *handler) GetLetterByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is letter by uuid"))
}

func (h *handler) UpdateLetter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is update letter"))
}

func (h *handler) PartiallyUpdateLetter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is partially update letter"))
}

func (h *handler) DeleteLetter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this delete letter"))
}
