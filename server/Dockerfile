# syntax=docker/dockerfile:1

# ---- Build Stage ----
FROM golang:1.19 AS builder

WORKDIR /app

# Кэшируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники сервера
COPY . .

# Сборка бинарника
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o server .

# ---- Final Stage ----
FROM gcr.io/distroless/base-debian11

WORKDIR /

# Копируем только бинарник
COPY --from=builder /app/server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/server"]