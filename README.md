# Blinkscan

## Building

### Build go backend
```sh
make build
```

### Run locally (requires running Postresql)
```sh
make run
```

### Generate new sqlc
```sh
make gen-sql
```

### Build new docker image
```sh
make image
```

### Run as docker compose cluster

First create a .env file according to .template.env in the root folder of this project 
```sh
docker compose up -d
```
