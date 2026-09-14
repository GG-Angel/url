# syntax=docker/dockerfile:1
# The microservice for shortening URLs.

# --- Build ---

FROM golang:1.27.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY internal/*.go ./internal/

RUN CGO_ENABLED=0 GOOS=linux go build -o /api

# --- Test ---

FROM builder AS tester
RUN go test -v ./...

# --- Run ---

FROM gcr.io/distroless/base-debian13:nonroot AS runner

COPY --from=builder /api /api
USER nonroot:nonroot

EXPOSE 8080
ENTRYPOINT ["/api"]
