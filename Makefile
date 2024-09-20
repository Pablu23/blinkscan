build: 
	go build -o bin/backend backend/cmd/blinkscan/backend/main.go

gen-sql:
	cd backend; sqlc generate

run: build
	./bin/backend
