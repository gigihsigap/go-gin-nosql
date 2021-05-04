# Golang Back-end

Using Go, Gin, and MongoDB.

## How to Run

Initialize project by:

1. Run `go mod tidy` to install dependencies
2. Make `.env` file to specify "MONGODB_URL"
3. Run the app by `go run main.go`
4. Connect to localhost:8000!

## Troubleshooting

### Postgres
> Collection name: "account"
> Database name: "test"
> Client name: "cluster0"

### SQLC
When running `make sqlc` your terminal might say "sqlc.json doesn't exist". You need to specify your project directory. Manually typing down the location works if $(pwd) somehow doesn't work.


## Features:

1. Basic Go, Gin, and MongoDB project structure.
2. API endpoints using Gin.
3. Validation of document entries.

This project is a bit lackluster due to my unfamiliarity with NoSQL projects.

## Endpoints

- **/accounts**

| Method | Header | Params | JSON                                                      |
| ------ | ------ | ------ | --------------------------------------------------------- |
| `POST` | `none` | `none` | username: `string`<br>line_id: `string` <br> email: `string` |
