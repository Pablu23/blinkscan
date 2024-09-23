VERSION = v0.0.1

build: 
	go build -o bin/blinkscan cmd/blinkscan/main.go

gen-sql:
	sqlc generate

run: build
	./bin/blinkscan

image: gen-sql
	docker build . -t pablu/blinkscan:$(VERSION) -t pablu/blinkscan:latest
