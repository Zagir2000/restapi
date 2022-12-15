package handlers

import (
	"github.com/julienschmidt/httprouter"
)

//Все наши хэндлеры будут реализовывать этот метод
type Handler interface {
	Register(router *httprouter.Router)
}
