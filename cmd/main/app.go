package main

import (
	"CRMka/internal/composites"
	"CRMka/internal/config"
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
	cfg := config.GetConfig()

	logger.Info("create mongodb composite")
	mongoDBComposite := composites.NewMongoDBComposite(context.Background(), *cfg)
	logger.Info("create employee composite")
	employeeComposite := composites.NewEmployeeComposite(logger, mongoDBComposite)

	logger.Info("create router")
	router := httprouter.New()
	logger.Infof("register handler employee handler")
	employeeComposite.Handler.Register(router)

	start(logger, router, cfg)
}

func start(logger *logging.Logger, router *httprouter.Router, cfg *config.Config) {
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
