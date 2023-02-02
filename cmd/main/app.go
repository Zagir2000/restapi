package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"restapi/internal/config"
	"restapi/internal/user"
	"restapi/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))

}

func Start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	//Запускаем как на айпи, так и с использованием сокета
	logger.Info("Start application")
	var listenErr error
	var listener net.Listener

	if cfg.Listen.Type == "sock" {
		//path/to/binary
		// После вызова метода Dir у нас останется /path/to
		fmt.Println(os.Args[0])
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		//Получим директорию в котором лежит приложение
		// Создадим сокет
		socketPath := path.Join(appDir, "app.sock")
		// создадим Unix socket
		logger.Debugf("socket path: %s", socketPath)
		logger.Info("Listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("Server is listening unix socket:%s", socketPath)
	} else {
		logger.Info("Listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.Host, cfg.Listen.Port))
		logger.Infof("Server is listening %s:%s", cfg.Listen.Host, cfg.Listen.Port)
	}
	if listenErr != nil {
		log.Fatal("Server off")
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatalln(server.Serve(listener))

}
func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()
	cfg := config.GetConfig()

	handler := user.NewHandler(logger)
	logger.Info("Register user handler")
	handler.Register(router)

	Start(router, cfg)
}
