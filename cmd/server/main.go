package main

import (
	"log"
	"time"

	"cbt/config"
	"cbt/database"
	"cbt/models"
	"cbt/routes"
	"cbt/utils"
	"gorm.io/gorm"
)

// updateAllExamTokens generates a new random token and updates it for all exams.
func updateAllExamTokens(db *gorm.DB) {
	newToken, err := utils.GenerateRandomToken(8) // Generate an 8-character random token
	if err != nil {
		log.Printf("Error generating new token: %v", err)
		return
	}

	// Update the token for all exams in a single query
	if err := db.Model(&models.Ujian{}).Where("1 = 1").Update("token", newToken).Error; err != nil {
		log.Printf("Error updating exam tokens: %v", err)
		return
	}

	log.Printf("Successfully updated exam token for all exams to: %s", newToken)
}

// startTokenUpdater starts a background goroutine to update exam tokens periodically.
func startTokenUpdater(db *gorm.DB) {
	log.Println("Starting background token updater...")
	ticker := time.NewTicker(5 * time.Minute)

	go func() {
		for {
			select {
			case <-ticker.C:
				updateAllExamTokens(db)
			}
		}
	}()
}

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate database schema
	log.Println("Running database migrations...")
	if err := db.AutoMigrate(
		&models.User{},
		&models.Siswa{},
		&models.Kelas{},
		&models.SiswaKelas{},
		&models.Ujian{},
		&models.UjianPesertaKelas{},
		&models.Soal{},
		&models.SoalPilihan{},
		&models.SesiUjian{},
		&models.JawabanSiswa{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Start the background token updater
	startTokenUpdater(db)

	// Setup router
	r := routes.SetupRouter(db, &cfg)

	// Start server
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
