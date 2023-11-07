package main

import (
	"CRMka/internal/company"
	"CRMka/internal/employee"
	"CRMka/internal/handlers"
	"CRMka/internal/letter"
	"CRMka/internal/project"
	"CRMka/internal/task"
	"CRMka/pkg/logging"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()

	logger.Info("create router")
	router := httprouter.New()

	myHandlers := []handlers.Handler{
		company.NewHandler(*logger),
		employee.NewHandler(*logger),
		letter.NewHandler(*logger),
		project.NewHandler(*logger),
		task.NewHandler(*logger),
	}

	for _, v := range myHandlers {
		v.Register(router)
		logger.Infof("handler %s registered", v.GetName())
	}

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	logger.Info("server is listenning port 0.0.0.0:1234")
	logger.Fatal(server.Serve(listener))
}
