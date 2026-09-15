# ============================
# Stage 1: Build stage
# ============================

FROM golang:1.24.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# Копируем go.work и его сумму
COPY go.work go.work.sum ./

# Копируем модули по папкам для кеширования зависимостей
COPY platform/go.mod platform/go.sum ./platform/
COPY shared/go.mod shared/go.sum ./shared/
COPY inventory/go.mod inventory/go.sum ./inventory/
COPY payment/go.mod payment/go.sum ./payment/
COPY order/go.mod order/go.sum ./order/

RUN go mod download

# Копируем весь исходный код
COPY platform ./platform
COPY shared ./shared
COPY inventory ./inventory
COPY order ./order
COPY payment ./payment

# Скачиваем grpc-health-probe
ADD https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.37/grpc_health_probe-linux-amd64 ./grpc-health-probe

RUN chmod +x grpc-health-probe

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app-inventory ./inventory/cmd/main.go

# ============================
# Stage 2: Final image
# ============================

FROM alpine:3.21.3

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Копируем бинарник и grpc-health-probe
COPY --from=builder /app/app-inventory .
COPY --from=builder /app/grpc-health-probe /bin/grpc-health-probe

COPY deploy/compose/inventory/.env deploy/compose/inventory/

EXPOSE 50051

# Запускаем под непривилегированным пользователем
USER appuser

ENTRYPOINT ["sh", "-c", "./app-inventory || (echo 'App crashed. Keeping container alive...' && sleep 3600)"]
