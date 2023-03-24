# MVC REST API and ORM Go Template

Have a template for starting a REST API with ORM integration and hot reload with go air.

## Video

https://yewtu.be/watch?v=ZI6HaPKHYsg

## REST API

Uses [Gin-Gonic](https://gin-gonic.com/docs/)

## ORM

Uses [GORM](https://gorm.io/)

## Hot Reload (air)

[Go Air](https://github.com/cosmtrek/air) enables hot reloading in go.
Usage of air:

1. Create of config file (.air.toml):
   - `air init`
2. Run app with hot reload (inside project root):
   - `air`
     Instead of:
   - `go run main.go`
   - `go run main.go server_config.go`

## Config YAML

Uses [Clean Env](https://github.com/ilyakaznacheev/cleanenv):

- go get -u github.com/ilyakaznacheev/cleanenv

## Swagger Docs (Swag)

Uses [Swag](https://github.com/swaggo/swag#how-to-use-it-with-gin)

- go install github.com/swaggo/swag/cmd/swag@latest
  Note: Go install tries to install the package into $GOBIN, when $GOBIN=/usr/local/go/bin will not work, works with $GOBIN=~/go/bin
  Initialize Swag (on project root):
- `swag init`
  Then get dependencies:
- `go get -u github.com/swaggo/gin-swagger`
- `go get -u github.com/swaggo/files`

Format Swag Comments:

- `swag fmt`

## TODO: Error Handling

## TODO: Understand router.Use(middlewares.AuthMiddleware())

## cURL Calls

- List users: GET /users:
  - `curl -X GET http://127.0.0.1:8080/users`
- Create user: POST /users:
  - `curl -X POST http://127.0.0.1:8080/users -H 'Content-Type: application/json' -d '{"name":"mario", "email":"mario@dhl.be", "password":"1234"}'`
- Get a user: GET /users/1:
  - `curl -X GET http://127.0.0.1:8080/users/1`
- Modify a user: PUT /users/1:
  - `curl -X PUT http://127.0.0.1:8080/users/1 -H 'Content-Type: application/json' -d '{"name":"john", "email":"john@dhl.be", "password":"5678"}'`
- Delete a user: DELETE /users/1:
  - `curl -X DELETE http://127.0.0.1:8080/users/1`
