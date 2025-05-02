FROM golang:1.24 AS base

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o totempedidos ./cmd/server/.

FROM alpine

WORKDIR /app

COPY --from=base /app/totempedidos .

CMD ["/app/totempedidos"]