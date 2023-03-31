# MVC REST API and ORM Go Template

Inspired by [go-gin-boilerplate](https://github.com/vsouza/go-gin-boilerplate)
for:

1. Swagger
2. Gin
3. Modular approach in go (no gorm dependency of controllers)

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

## JWT Authentication

Tutorial: [https://codewithmukesh.com/blog/jwt-authentication-in-golang/](https://codewithmukesh.com/blog/jwt-authentication-in-golang/)
Check JWT token [decoder](https://jwt.io/)
Dependencies:

- go get github.com/dgrijalva/jwt-go (deprecated)
- go get -u github.com/golang-jwt/jwt/v5 (up to date)
- go get golang.org/x/crypto/bcrypt

1. Authenticate (create JWT token for the user account):

- `curl -X POST http://127.0.0.1:8080/api/v1/token -d '{"email":"alice@gmail.com", "password":"1234"}'`
  Response:
  `{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlODkiLCJlbWFpbCI6ImFsaWNlQGdtYWlsLmNvbSIsImV4cCI6MTY4MDI2MDU2OH0.6d9-WiCQTAcs4wxIkqHyQ3J0-UZBEr2_swpdUcO7zRc"
}`

2. Authorized Call:

- `curl -X GET http://127.0.0.1:8080/api/v1/secured/ping -H 'Accept: application/json' -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlODkiLCJlbWFpbCI6ImFsaWNlQGdtYWlsLmNvbSIsImV4cCI6MTY4MDI2MDU2OH0.6d9-WiCQTAcs4wxIkqHyQ3J0-UZBEr2_swpdUcO7zRc'`

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

Swagger API:

- `http://localhost:8080/swagger/index.html#/`

## Error Handling

Inspired by:

- https://blog.depa.do/post/gin-validation-errors-handling

### Logging

Zero allocation JSON logger, [zerolog](https://github.com/rs/zerolog):

- `go get -u github.com/rs/zerolog/log`

### Go Validator

Gin uses [Go Validator v10] (https://github.com/go-playground/validator):

- `go get github.com/go-playground/validator/v10`

In code import:

- `import "github.com/go-playground/validator/v10" `

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
