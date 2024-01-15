build: 
	@templ generate
	@go build -o bin/kagami cmd/http/main.go

run: build
	@./bin/kagami

seed: 
	@go run cmd/seed/seed.go
