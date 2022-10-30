package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (c chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (c chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (c chiRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Put(uri, f)
}

func (c chiRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Delete(uri, f)
}

func (c chiRouter) SERVE(port string) {
	log.Fatal(http.ListenAndServe(port, chiDispatcher))
}
