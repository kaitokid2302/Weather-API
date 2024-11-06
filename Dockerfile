from golang:1.23-alpine

workdir /app

copy . .

run go mod tidy

cmd ["go", "run", "./cmd/main.go"]
