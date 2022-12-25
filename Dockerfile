# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

WORKDIR /

RUN apk update 
RUN apk add --no-cache gcc g++ git openssh-client

COPY . .

RUN go mod download

RUN GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn,direct \
    go build -ldflags="-extldflags=-static" -tags sqlite_omit_load_extension -o ./main ./main.go

# Build the server image

# gcr.io/distroless/base
FROM gcr.io/distroless/static-debian11

WORKDIR /root/

COPY --from=build /main .
COPY --from=build ./config ./config

CMD ["./main"]