VERSION = v0.0.1

build: 
	go build -o bin/backend backend/cmd/blinkscan/backend/main.go

gen-sql:
	cd backend; sqlc generate

run: build
	./bin/backend

docker-image: gen-sql
	sudo docker build ./backend/ -t pablu/blinkscan-backend:$(VERSION)
