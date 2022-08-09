# go_projects

## Inspirations

- https://www.youtube.com/watch?v=JzdPOq2kgfk

## First_api

- go mod init api/server
- go get github.com/gin-gonic/gin
- Terminal 1: go run main.go && Terminal 2: curl localhost:6666/markets

-- curl localhost:6666/markets --header "Content-Type: application/json" -d @body.json --request POST
