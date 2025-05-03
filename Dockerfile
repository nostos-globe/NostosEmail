FROM golang:1.24 AS builder
 
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o profile-service ./cmd/main.go

# Imagen final para producción (más ligera)
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copia el binario compilado desde el builder
COPY --from=builder /app/email-service .

# Ejecuta el servicio
CMD ["/app/email-service"]
