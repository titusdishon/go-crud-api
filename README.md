# A Template For Implementing Clean Architecture Using golang, Mysql and  Docker

In this repo I have added the necessary configurations for containerizing a golang application.
I have also followed the clean architecture pattern where there is no dependency on:

- No database dependency
- Http client

This separates data storage logic from business logic.
I have implemented two different http clients, gorilla mux and go-chi.
To switch between these clients change directory into routes directory and open the file routes.go.
Change the following lines accordingly.

``` go
var (
 repo       repositories.UserRepository = repositories.NewMysqlRepository()
 service    services.UserService        = services.NewUserService(repo)
 // httpRouter router.Router               = router.NewMuxRouter()   <---- change this
 httpRouter router.Router               = router.NewChiRouter()  // <--- to this
 controller controllers.IUserController = controllers.NewUserController(service)
)

```

## Tech stack

- Mysql
- Golang
- Docker & docker-compose

## Project Structure

``` bash

.
├── README.md
├── api.Dockerfile // api dockerfile
├── config     // Connection to mysql database
│   └── mysql_connect.go
├── controllers    // provides access to the services
│   └── user_controllers.go
├── database      // initial database structure
│   └── migration.sql
├── db.Dockerfile // database dockerfile
├── docker-compose.yml  // automate docker process
├── entity   // define data structures for entities
│   └── user.go
├── entrypoint.sh   
├── errors   // error formatter 
│   └── errors.go
├── go.mod
├── go.sum
├── http   // defines the http client, new clients can be added
│   ├── chi-router.go
│   ├── mux-router.go
│   └── routes.go
├── main
├── main.go   // application entry point
├── main_test.go
├── repositories  // provides access to the database, ways to carry out CRUD operations
│   └── mysql_repositories.go
├── routes     // defines routes for the application
│   └── routes.go
├── services  // provides access to services 
│   ├── user-service.go
│   └── user_service_test.go
├── swagger-docs  // swagger documentation
│   ├── checkapi.yaml
│   ├── openapi.yaml
│   ├── resources
│   │   ├── check.yaml
│   │   └── user.yaml
│   ├── responses  // all responses
│   │   ├── bad_request.yaml
│   │   ├── not_found.yaml
│   │   └── success.yaml
│   └── schemas // all schemas, fields to be accessed 
│       ├── security_schema.yaml
│       └── user_schema.yaml
└── utils   // common functions to be re-used
    └── utils.go

14 directories, 32 files
```

### Running The project on your local setup

- Clone the project
- Change directory into the project root directory on the terminal
- Create .env file and add the following code:

```.env

MYSQL_RANDOM_ROOT_PASSWORD: "secret"
MYSQL_DATABASE: "test_database"
MYSQL_USER: "test_user"
MYSQL_PASSWORD: "secret"
MYSQL_PORT: 3306
SWAGGER_PORT: 8080
MYSQL_HOST:db
APP_PORT:8081

```

- If you have docker installed and configured in your local machine
- To build the project:

 ```.env
docker compose build --no-cache
 ```

- To run the project run:

```.env
docker compose up
 ```

- The project API should now be available in your local setup using the url [http:localhost:8081](http:localhost:8081)
- The project Swagger documentation should now be available in your local setup using the url [http:localhost:8080](http:localhost:8080)
