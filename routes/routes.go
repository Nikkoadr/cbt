package routes

import (
	"cbt/config"
	"cbt/handlers"
	"cbt/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Handlers
	examHandler := &handlers.ExamHandler{DB: db}
	adminHandler := &handlers.AdminHandler{DB: db}

	// Root group
	api := r.Group("/api")

	// Public route for login
	api.POST("/login", middleware.LoginHandler(cfg.JWTSecret, db))

	// Routes for Siswa, protected by SiswaMiddleware
	siswaRoutes := api.Group("/siswa")
	siswaRoutes.Use(middleware.SiswaMiddleware(cfg.JWTSecret))
	{
		siswaRoutes.GET("/ujian/daftar", examHandler.GetDaftarUjian)
		siswaRoutes.POST("/ujian/:id/mulai", examHandler.MulaiUjian)
		siswaRoutes.POST("/ujian/sesi/:sesiID/simpan", examHandler.SimpanJawaban)
		siswaRoutes.POST("/ujian/sesi/:sesiID/selesai", examHandler.SelesaikanUjian)
	}

	// Routes for Admin/Pengawas, protected by AdminMiddleware
	adminRoutes := api.Group("/admin")
	adminRoutes.Use(middleware.AdminMiddleware(cfg.JWTSecret))
	{
		adminRoutes.POST("/ujian/sesi/:sesiID/koreksi", adminHandler.KoreksiUjian)
		adminRoutes.GET("/ujian/:id/token", adminHandler.GetUjianToken)
	}

	return r
}
