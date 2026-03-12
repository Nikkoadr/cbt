package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/GoogleCloudPlatform/golang-samples/run/helloworld/config"
	"github.com/GoogleCloudPlatform/golang-samples/run/helloworld/handlers"
	"github.com/GoogleCloudPlatform/golang-samples/run/helloworld/middleware"
)

// SetupRouter sets up the router
func SetupRouter(db *gorm.DB, cfg config.Config) *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Adjust for your frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Handlers
	authHandler := &handlers.AuthHandler{DB: db, Cfg: cfg}
	examHandler := &handlers.ExamHandler{DB: db}

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API sudah berjalan"})
	})

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", authHandler.Logout)
			auth.GET("/me", middleware.AuthMiddleware(cfg), authHandler.Me)
		}

		exam := api.Group("/exam")
		exam.Use(middleware.AuthMiddleware(cfg))
		{
			exam.GET("/today", examHandler.GetTodayExams)
			exam.GET("/:id", examHandler.GetExamDetails)
			exam.GET("/:id/access", examHandler.CheckExamAccess)
		}
	}

	return r
}
