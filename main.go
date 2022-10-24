package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/titusdishon/go-docker-mysql/routes"
)

func main() {
	routes.UserRouters()

}
