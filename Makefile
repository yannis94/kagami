build: 
	@templ generate
	@go build -o bin/kagami cmd/main.go

run: build
	@./bin/kagami
