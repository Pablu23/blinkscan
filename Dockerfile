# Debug
# FROM busybox:uclibc AS busybox
FROM node:alpine AS ng-build-stage
WORKDIR /app

RUN npm install -g @angular/cli
COPY ./frontend/package.json ./frontend/package-lock.json ./
RUN npm install

COPY ./frontend .

RUN ng build
RUN rm -rf ./frontend


FROM golang:1.23 AS go-build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-blinkscan ./cmd/blinkscan/main.go


FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /app

COPY --from=go-build-stage /app/docker-blinkscan . 
COPY --from=ng-build-stage /app/dist/blinkscan ./frontend

# Debug
# COPY --from=busybox /bin/sh /bin/sh
# COPY --from=busybox /bin/cat /bin/cat
# COPY --from=busybox /bin/ls /bin/ls

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/docker-blinkscan"]
