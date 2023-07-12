# MVC REST API and ORM Go Template

Inspired by [go-gin-boilerplate](https://github.com/vsouza/go-gin-boilerplate)
for:

1. Swagger
2. Gin
3. Modular approach in go (no gorm dependency of controllers)

Have a template for starting a REST API with ORM integration and hot reload with go air.

## Run Environments

Setup 3 environments: dev, stage and prod

### Run Dev Environment

#### Run DB (terminal 1)

1. cd docker/dev
2. docker compose up

##### Enter PostgreSQL DB (terminal 2)

1. `docker exec -it dev-db-1 bash`
2. `su - postgres`
3. `psql -h localhost -p 5432 -U <user> <db_name>`
4. `psql -h localhost -p 5432 -U postgres postgres`
5. `\l` : list DBs
6. `\c <db_name>` : connect to DB
7. `\dt` : list all the tables
8. `select \* from <table_name>;`
9. `select \* from users;`

#### Run Go REST API Service (terminal 3)

1. `cd <project_root>`
2. `SERVER_ENVIRONMENT="dev" go run main.go`

## REST API

Uses [Gin-Gonic](https://gin-gonic.com/docs/)

## ORM

Uses [GORM](https://gorm.io/)

## Format Code

1. `cd ~/go/src/github.com/pavva91/gin-gorm-rest/`
2. `gofmt -l -s -w .`

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

- `go install github.com/swaggo/swag/cmd/swag@latest`
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

## Unit Tests

- `go get -u github.com/stretchr/testify`
- Run all tests: `go test ./...`
- Run all tests and create code coverage report: `go test -v -coverprofile cover.out ./...`
- Run all tests with code coverage and open on browser: `go test -v -coverprofile cover.out ./... && go tool cover -html=cover.out`
- Run tests of /controllers with code coverage and open on browser: `go test -v -coverprofile cover.out ./controllers/ && go tool cover -html=cover.out`
- Run tests of /services with code coverage and open on browser: `go test -v -coverprofile cover.out ./services/ && go tool cover -html=cover.out`
- Run tests of /repositories with code coverage and open on browser: `go test -v -coverprofile cover.out ./repositories/ && go tool cover -html=cover.out`
- Run specific tests (regex): `go test -run TestMyFunction ./...`
- Run specific tests (regex): `go test -run Test_GenerateToken_InvalidRequestBodyNoPasswordField_400BadRequest ./...`
- Run specific tests (regex) of a module (e.g. controllers): `go test -run Test_GenerateToken ./controllers`
- Run tests of a folder (e.g. validation): `go test ./validation`
- Run tests of a module (e.g. validation): `go test github.com/pavva91/gin-gorm-rest/validation`
- Run particular test (e.g. test Test_GetByID_EmptyId_400BadRequest in /controllers/user_test.go): `go test -v ./... -run Test_GetByID_EmptyId_400BadRequest`
- Run particular test (e.g. test Test_GetByID_EmptyId_400BadRequest in /controllers/user_test.go): `go test -v ./controllers -run Test_GetByID_EmptyId_400BadRequest`
- Run particular test with code coverage report (e.g. test Test_GetByID_EmptyId_400BadRequest in /controllers/user_test.go): `go test -v -coverprofile cover.out ./controllers -run Test_GetByID_EmptyId_400BadRequest && go tool cover -html=cover.out`
- Run particular test when using test suite (e.g. test Test_Status_Return200 in TestSuiteHealthController in /controllers/health_test.go): `go test -v ./controllers -run TestSuiteHealthController -testify.m Test_Status_Return200`

### Code Structure to enable Unit Tests (testeable code)

- [Code samples](https://github.com/federicoleon/golang-examples/tree/master/testeable_code)

### Code Coverage

- By package name:

  - Just run: `go test -cover github.com/pavva91/gin-gorm-rest/validation`
  - Create coverage file: `go test -v -coverprofile cover.out github.com/pavva91/gin-gorm-rest/validation`
  - One Command create coverage file and open in browser; `go test -v -coverprofile cover.out github.com/pavva91/gin-gorm-rest/controllers && go tool cover -html=cover.out`

- By folder:

  - Just run: `go test -cover ./validation`
  - Create coverage file: `go test -v -coverprofile cover.out ./validation`
  - Open coverage file on browser: `go tool cover -html=cover.out`
  - One Command create coverage file and open in browser; `go test -v -coverprofile cover.out ./controllers/ && go tool cover -html=cover.out`

- Run all tests:
  - Just run: `go test ./... -cover`
  - Create coverage file: `go test ./... -coverprofile coverage.out`
  - Open coverage file on browser: `go tool cover -html=coverage.out`

From [stack overflow](https://stackoverflow.com/questions/10516662/how-to-measure-test-coverage-in-go)

1. Create function in ~/.bashrc and/or ~/.zshrc:

```bash
cover () {
  t="/tmp/go-cover.$$.tmp"
  go test -coverprofile=$t $@ && go tool cover -html=$t && unlink $t
}
```

2. Call this function:

- `cd ~/go/src/github.com/pavva91/gin-gorm-rest/ `
- `cover github.com/pavva91/gin-gorm-rest/validation`

### Run Unit Test with Debugger

- :lua require('dap-go').debug_test()
- Keymap: <leader>dt

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
