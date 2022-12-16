package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"restapi/internal/user"
	"restapi/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))

}

func Start(router *httprouter.Router) {
	logger := logging.GetLogger()
	//Запускаем как на айпи, так и с использованием сокета
	logger.Info("Start application")
	listenet, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Server off")
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Info("Server is listening 127.0.0.1:1234")
	logger.Fatalln(server.Serve(listenet))

}
func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()
	handler := user.NewHandler(logger)
	logger.Info("Register user handler")
	handler.Register(router)
	Start(router)
}
