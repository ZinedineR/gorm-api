package http

import (
	"github.com/labstack/echo/v4"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.
func NewGinEngineTV(tvHandler *TVHandler, internalUsername, internalPassword string) *echo.Echo {
	engine := echo.New()

	// CORS
	// engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowAllOrigins: true,
	// 	AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization", "SVC_USER", "SVC_PASS"},
	// 	AllowMethods:    []string{"GET", "POST", "PUT", "PATCH"},
	// }))

	engine.GET("/", Status)
	// engine.GET("/healthz", Health)
	engine.GET("/version", Version)

	engine.POST("/create-TV", tvHandler.CreateTV)
	engine.GET("/list-TV", tvHandler.GetListTV)
	engine.GET("/get-TV/:id", tvHandler.GetDetailTV)
	engine.PUT("/update-TV/:id", tvHandler.UpdateTV)
	engine.DELETE("/delete-TV/:id", tvHandler.DeleteTV)

	return engine
}
