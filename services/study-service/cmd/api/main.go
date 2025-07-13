package main

import (
	"context"
	"log"
	"study-service/internal/handler"
	"study-service/internal/repository"
	"study-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// 1. Konfigürasyon Yükleme (Placeholder)
	// Normalde bu kısımda .env dosyasından veya ortam değişkenlerinden
	// veritabanı bağlantı bilgisi gibi konfigürasyonlar okunur.
	// dbURL := os.Getenv("DATABASE_URL")
	dbURL := "postgres://user:password@localhost:5432/study_db" // Örnek bağlantı adresi

	// 2. Veritabanı Bağlantısı (Placeholder)
	// Gerçek bir uygulamada, uygulama kapanırken db.Close() çağrılmalıdır.
	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v\n", err)
	}
	defer db.Close()
	log.Println("Veritabanı bağlantısı başarılı.")

	// 3. Katmanları Oluşturma (Dependency Injection)
	noteRepo := repository.NewPostgresNoteRepository(db)
	noteService := service.NewNoteService(noteRepo)
	noteHandler := handler.NewNoteHandler(noteService)

	// 4. Gin Router'ı Başlatma
	router := gin.Default()

	// 5. Rotaları Tanımlama
	router.POST("/api/v1/notes", noteHandler.CreateNote)

	// 6. Sunucuyu Başlatma
	log.Println("Sunucu http://localhost:8080 adresinde başlatılıyor...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Sunucu başlatılamadı: %v", err)
	}
}
