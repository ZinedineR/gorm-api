package http

import (
	"github.com/labstack/echo/v4"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.
func NewGinEngine(detailedHandler *DetailedHandler, actorHandler *ActorHandler, internalUsername, internalPassword string) *echo.Echo {
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
	//Detailed
	engine.POST("/create-Detailed", detailedHandler.CreateDetailed)
	engine.GET("/list-Detailed", detailedHandler.GetListDetailed)
	engine.GET("/get-Detailed/:id", detailedHandler.GetDetailDetailed)
	engine.PUT("/update-Detailed/:id", detailedHandler.UpdateDetailed)
	engine.DELETE("/delete-Detailed/:id", detailedHandler.DeleteDetailed)
	//Actor
	engine.POST("/create-Actor", actorHandler.CreateActor)
	engine.GET("/list-Actor", actorHandler.GetListActor)
	engine.GET("/get-Detailed/:id", actorHandler.GetDetailActor)
	engine.PUT("/update-Detailed/:id", actorHandler.UpdateActor)
	engine.DELETE("/delete-Detailed/:id", actorHandler.DeleteActor)
	return engine
}
