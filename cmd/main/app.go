package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"restapi/internal/user"
	"time"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))

}

func Start(router *httprouter.Router) {
	//Запускаем как на айпи, так и с использованием сокета
	log.Println("Start aplication")
	listenet, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Server off")
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	server.Serve(listenet)
	log.Println("Server is listening 127.0.0.1:1234")
}
func main() {
	log.Println("Create router")
	router := httprouter.New()
	handler := user.NewHandler()
	log.Println("Register user handler")
	handler.Register(router)
	Start(router)
}
