package http

import (
	"github.com/labstack/echo/v4"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.
func NewGinEngine(tvHandler *TVHandler, streamedHandler *StreamedHandler, watchedHandler *WatchedHandler, detailedHandler *DetailedHandler, actorHandler *ActorHandler, internalUsername, internalPassword string) *echo.Echo {

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

	//tv
	engine.POST("/create-TV", tvHandler.CreateTV)
	engine.GET("/list-TV", tvHandler.GetListTV)
	engine.GET("/get-TV/:id", tvHandler.GetDetailTV)
	engine.PUT("/update-TV/:id", tvHandler.UpdateTV)
	engine.DELETE("/delete-TV/:id", tvHandler.DeleteTV)
	//streamed
	engine.POST("/create-streamed", streamedHandler.CreateStreamed)
	engine.GET("/list-streamed", streamedHandler.GetListStreamed)
	engine.GET("/get-streamed/:id", streamedHandler.GetDetailStreamed)
	engine.PUT("/update-streamed/:id", streamedHandler.UpdateStreamed)
	engine.DELETE("/delete-streamed/:id", streamedHandler.DeleteStreamed)
	//watched
	engine.POST("/create-watched", watchedHandler.CreateWatched)
	engine.GET("/list-watched", watchedHandler.GetListWatched)
	engine.GET("/get-watched/:id", watchedHandler.GetDetailWatched)
	engine.PUT("/update-watched/:id", watchedHandler.UpdateWatched)
	engine.DELETE("/delete-watched/:id", watchedHandler.DeleteWatched)
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
