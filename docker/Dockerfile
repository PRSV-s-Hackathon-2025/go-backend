# syntax=docker/dockerfile:1

# Stage 1 — build
FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server

# Stage 2 — run
FROM gcr.io/distroless/base-debian11

WORKDIR /app
COPY --from=builder /app/server /app/server
COPY environment /app/environment

EXPOSE 8002

ENTRYPOINT ["/app/server"]
