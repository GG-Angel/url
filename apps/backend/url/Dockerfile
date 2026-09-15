# syntax=docker/dockerfile:1
FROM golang:1.27.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /url ./cmd

FROM builder AS tester
RUN go test -v ./...

FROM gcr.io/distroless/base-debian13:nonroot AS runner

COPY --from=builder /url /url
USER nonroot:nonroot

EXPOSE 8080
ENTRYPOINT ["/url"]
