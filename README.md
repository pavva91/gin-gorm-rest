# MVC REST API and ORM Go Template

Have a template for starting a REST API with ORM integration and hot reload with go air.

## Video

https://yewtu.be/watch?v=ZI6HaPKHYsg

## REST API

Uses [Gin-Gonic](https://gin-gonic.com/docs/)

## ORM

Uses [GORM](https://gorm.io/)

## Hot Reload (air)

Usage of air:

1. Create of config file (.air.toml):
   - `air init`
2. Run app with hot reload (inside project root):
   - `air`
     Instead of:
   - `go run main.go`

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
