# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

RUN apk update 
RUN apk add --no-cache gcc g++ git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn,direct \
    go build -ldflags="-extldflags=-static" -o ./main ./main.go

# Build the server image

# gcr.io/distroless/base
FROM gcr.io/distroless/static-debian11

WORKDIR /root/

COPY --from=build /app/config ./config
COPY --from=build /app/main .

CMD ["./main"]