FROM golang:1.26 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" .

FROM scratch
COPY --from=builder /app/ftpput /ftpput
ENTRYPOINT ["/ftpput"]
