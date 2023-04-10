FROM golang:1.18-alpine AS builder
WORKDIR /app 
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
COPY initializers ./initializers/
COPY models ./models/
COPY migrate ./migrate/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migrate ./migrate/migrate.go

FROM golang:1.18-alpine
COPY --from=builder /app/migrate /usr/bin/
CMD ["/usr/bin/migrate"]
