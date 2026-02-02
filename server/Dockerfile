# syntax=docker/dockerfile:1

# ---- Build Stage ----
FROM golang:1.19 AS builder

WORKDIR /app

# Кэшируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники сервера
COPY . .

RUN addgroup --gid 1001 appuser && \
    adduser --disabled-password --uid 1001 --ingroup appuser appuser

# Сборка бинарника
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o server .

# ---- Final Stage ----
FROM gcr.io/distroless/base-debian11

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group


WORKDIR /

# Копируем только бинарник
COPY --from=builder /app/server /server

USER appuser

EXPOSE 8080

ENTRYPOINT ["/server"]