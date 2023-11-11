package main

import (
	"CRMka/internal/company"
	"CRMka/internal/config"
	"CRMka/internal/employee"
	"CRMka/internal/handlers"
	"CRMka/internal/letter"
	"CRMka/internal/project"
	"CRMka/internal/task"
	"CRMka/pkg/client/mongodb"
	"CRMka/pkg/logging"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	cfgMongo := cfg.MongoDB
	mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username,
		cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	if err != nil {
		panic(err)
	}
	//1 storageEmplooye := db.NewStorage(mongoDBClient, "employees", logger)

	SetHandlers := []handlers.Handler{
		company.NewHandler(logger),
		employee.NewHandler(logger),
		letter.NewHandler(logger),
		project.NewHandler(logger),
		task.NewHandler(logger),
	}

	for _, v := range SetHandlers {
		logger.Infof("register handler %s", v.GetName())
		v.Register(router)
	}

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error
	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket: %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
