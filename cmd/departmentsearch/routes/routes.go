package routes

import (
	"braiton/braiton-home/cmd/departmentsearch/routes/handlers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes() {
	port := os.Getenv("PORT")

	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "5000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("api/entidad", handlers.GetEntities)
	router.GET("api/entidad/:id", handlers.GetEntity)
	router.POST("api/entidad", handlers.CreateEntity)
	router.PUT("api/entidad/:id", handlers.UpdateEntity)
	router.DELETE("api/entidad/:id", handlers.DeleteEntity)

	router.Run(":" + port)
}
