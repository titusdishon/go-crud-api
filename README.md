# Containerizing Golang applications with docker using docker compose

In this repo I have added the necessary configurations for containerizing a golang application for m1 users

## Tech stack

- Mysql
- Golang
- Docker & docker-compose

### Running The project on your local setup

- Clone the project
- Change directory into the project root directory on the terminal
- Create .env file and add the following code:

```.env

MYSQL_RANDOM_ROOT_PASSWORD: "some_secret"
MYSQL_DATABASE: "database_name"
MYSQL_USER: "some_user"
MYSQL_PASSWORD: "some_secret"
PORT:8080

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

- The project should now be available in your local setup using the url [http:localhost:8080](http:localhost:8080)
