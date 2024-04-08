# User Microservice

This User Microservice is a RESTful API developed with the Go programming language, utilizing the Gin framework for routing, GORM for object-relational mapping, and PostgreSQL as the database backend. It supports basic CRUD operations for managing user data.

### Prerequisites

- Go
- Docker

### Running the Service

To run the service with docker-compose
```shell
docker-compose up --build
```
The service will start and listen on port 8080.


To run the service without docker-compose
Firstly you need to have a postgres database running on your local machine.

```shell
docker run --name mypostgres -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=user_microservice -p 5432:5432 -d postgres
```

Then you can run the service with the following command
```shell
go run main.go
```

or with built binary for MacOS system

```shell
./user-microservice
```

The service will start and listen on port 8080.


## API Endpoints

Below are the available API endpoints and example `curl` commands to interact with them.

### Create User

- **POST** `/save`

Creates a new user.


```shell
curl -X POST http://localhost:8080/save \
-d '{"name": "John Doe", "email": "john.doe@example.com", "date_of_birth": "1990-01-01"}'
```

### Get User by ID

- **GET** `/:id`

Retrieves a user by their ID.


```shell
curl -X GET http://localhost:8080/1
```

### Update User

- **PUT** `/:id`

Updates an existing user.


```shell
curl -X PUT http://localhost:8080/1 \
-d '{"name": "Jane Doe", "email": "jane.doe@example.com", "date_of_birth": "1991-02-01"}'
```

### Delete User

- **DELETE** `/:id`

Deletes a user by their ID.


```shell
curl -X DELETE http://localhost:8080/1
```

### Get All Users

- **GET** `/`

Retrieves all users.


```shell
curl -X GET http://localhost:8080/
```

### Delete All Users

- **DELETE** `/`

Deletes all users.


```shell
curl -X DELETE http://localhost:8080/
```
