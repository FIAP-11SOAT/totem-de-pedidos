ARG GO_VERSION=1.24

# Build
FROM golang:${GO_VERSION}-alpine AS build
WORKDIR /service
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app ./service/cmd/worker

# Test
FROM golang:${GO_VERSION}-alpine AS tests
ENV CI "1"
WORKDIR /service
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN go clean -testcache
CMD go test -v ./test/unittests/...


# Image
FROM gcr.io/distroless/static-debian12 AS production
USER nonroot:nonroot
COPY --from=build --chown=nonroot:nonroot /app /app
ENTRYPOINT ["/app"]