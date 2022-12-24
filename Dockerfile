# syntax=docker/dockerfile:1

# Build a golang image based on https://docs.docker.com/language/golang/build-images
FROM golang:1.18-alpine AS build

WORKDIR /

RUN apk update 
RUN apk add --no-cache gcc g++ git openssh-client

COPY . .

RUN go mod download

RUN GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn,direct \
    go build -ldflags="-extldflags=-static" -tags sqlite_omit_load_extension -o ./server ./cmd/server/server.go

# Build the server image
FROM gcr.io/distroless/base

# RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /server .
COPY --from=build ./database/test.db ./database/test.db

EXPOSE 8080

CMD ["./server"]