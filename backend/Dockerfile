FROM golang:1.23 AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-backend ./cmd/blinkscan/main.go

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /app

COPY --from=build-stage /app/docker-backend . 

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/docker-backend"]
