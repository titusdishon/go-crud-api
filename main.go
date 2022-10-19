package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/titusdishon/go-docker-mysql/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load env files")
	}
	r := mux.NewRouter()
	routes.UserRouters(r)
	http.Handle("/", r)
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(PORT, r))
}
