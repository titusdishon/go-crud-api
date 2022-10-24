package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

var (
	muxDispatcher = mux.NewRouter()
)

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}
func (*muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("PUT")
}
func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}
func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux running on port: %s", port)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	http.Handle("/", muxDispatcher)
	handler := c.Handler(muxDispatcher)
	log.Fatal(http.ListenAndServe(port, handler))
}
