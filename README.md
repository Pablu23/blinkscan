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

### Run as docker compose cluster
```sh
docker compose up -d
```

### Generate new sqlc
```sh
make gen-sql
```

### Build new docker image
```sh
make image
```
