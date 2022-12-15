package user

import (
	"net/http"
	"restapi/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

const (
	usersURl = "/users"
	userURl  = "/users/:uuid"
)

type handler struct {
}

//создаем структуру, но возвращаем интерфейс
func NewHandler() handlers.Handler {
	return &handler{}
}

//Наша структура handler реализует интерфейс хэндлер
func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURl, h.GetUsers)
	router.GET(userURl, h.GetUserByUUID)
	router.POST(userURl, h.CreateUser)
	router.PUT(userURl, h.UpdateUser)
	router.PATCH(userURl, h.PartiallyUpdate)
	router.DELETE(userURl, h.DeleteUser)
}
func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("This is list users"))
}
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("This is user by UUID"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("This is update user"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("This is update user"))
}
func (h *handler) PartiallyUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("This is partially update user"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is delete user"))
	w.WriteHeader(204)
}
