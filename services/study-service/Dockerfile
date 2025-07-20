# --- Aşama 1: Builder ---
# Bu aşama, Go kaynak kodunu derlemek için kullanılır.
FROM golang:1.22-alpine AS builder

# Güvenlik güncellemelerini al
RUN apk update && apk upgrade

# Uygulama için çalışma dizini oluştur
WORKDIR /app

# Bağımlılıkları önbelleğe almak için önce modül dosyalarını kopyala
COPY go.mod go.sum ./
RUN go mod download

# Kaynak kodunun geri kalanını kopyala
COPY . .

# Uygulamayı derle. CGO_ENABLED=0 ile statik bir binary oluşturulur.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./study-service ./cmd/api/main.go

# --- Aşama 2: Final ---
# Bu aşama, sadece derlenmiş binary'yi içeren küçük bir imaj oluşturur.
FROM alpine:latest

WORKDIR /app

# Derlenmiş binary'yi builder aşamasından kopyala
COPY --from=builder /app/study-service .

# Uygulamanın çalışacağı portu belirt
EXPOSE 8080

# Container başladığında çalıştırılacak komut
CMD ["./study-service"]