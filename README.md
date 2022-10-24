# Implementing Clean Architecture Using golang, Mysql and  Docker

In this repo I have added the necessary configurations for containerizing a golang application.
I have also followed the clean architecture pattern where there is no dependency on:

- No database dependency
- Http library
Thus separating the business logic from data storage logic.
This way you can easily switch from one database to the other without needing much configuration

## Tech stack

- Mysql
- Golang
- Docker & docker-compose

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

- The project should now be available in your local setup using the url [http:localhost:8081](http:localhost:8081)
