# argument for Go version
ARG GO_VERSION=1.19.2
 
# STAGE 1: building the executable
FROM golang:${GO_VERSION}-alpine AS build
 
RUN apk add --no-cache git
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
# Run tests
# RUN CGO_ENABLED=0 go test -timeout 30s -v github.com/gbaeke/go-template/pkg/api
# Build the executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -installsuffix 'static' \
    -ldflags="-w -s" \
    -o /app ./cmd/app
 
# STAGE 2: build the container to run
FROM gcr.io/distroless/static AS final
 
LABEL maintainer="runior"
USER nonroot:nonroot
 
# copy compiled app
COPY --from=build --chown=nonroot:nonroot /app /app
 
# run binary; use vector form
ENTRYPOINT ["/app"]