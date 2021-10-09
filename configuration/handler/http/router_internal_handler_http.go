package http

import (
	"github.com/labstack/echo"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.
func NewGinEngine(tvHandler *TVHandler, streamedHandler *StreamedHandler, watchedHandler *WatchedHandler, detailedHandler *DetailedHandler, actorHandler *ActorHandler, h *Loginhandler, internalUsername, internalPassword string) *echo.Echo {

	engine := echo.New()

	// CORS
	// engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowAllOrigins: true,
	// 	AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization", "SVC_USER", "SVC_PASS"},
	// 	AllowMethods:    []string{"GET", "POST", "PUT", "PATCH"},
	// }))

	engine.GET("/", Status)
	engine.POST("/login", h.Login)
	engine.GET("/private", h.Private, IsLoggedIn)
	engine.GET("/admin", h.Private, IsLoggedIn, isAdmin)
	engine.GET("/version", Version)

	//tv
	engine.POST("/create-TV", tvHandler.CreateTV, IsLoggedIn, isAdmin)
	engine.GET("/list-TV", tvHandler.GetListTV, IsLoggedIn)
	engine.GET("/get-TV/:id", tvHandler.GetDetailTV, IsLoggedIn)
	engine.PUT("/update-TV/:id", tvHandler.UpdateTV, IsLoggedIn, isAdmin)
	engine.DELETE("/delete-TV/:id", tvHandler.DeleteTV, IsLoggedIn, isAdmin)
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
	engine.POST("/create-detailed", detailedHandler.CreateDetailed)
	engine.GET("/list-detailed", detailedHandler.GetListDetailed)
	engine.GET("/get-detailed/:id", detailedHandler.GetDetailDetailed)
	engine.PUT("/update-detailed/:id", detailedHandler.UpdateDetailed)
	engine.DELETE("/delete-detailed/:id", detailedHandler.DeleteDetailed)
	//Actor
	engine.POST("/create-actor", actorHandler.CreateActor)
	engine.GET("/list-actor", actorHandler.GetListActor)
	engine.GET("/get-actor/:id", actorHandler.GetDetailActor)
	engine.PUT("/update-actor/:id", actorHandler.UpdateActor)
	engine.DELETE("/delete-actor/:id", actorHandler.DeleteActor)
	//User
	engine.POST("/create-user", userHandler.CreateUser)

	return engine
}
