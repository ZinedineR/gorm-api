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

func NewGinEngineStreamed(streamedHandler *StreamedHandler, internalUsername, internalPassword string) *echo.Echo {
	engine := echo.New()

	// CORS
	// engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowAllOrigins: true,
	// 	AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization", "SVC_USER", "SVC_PASS"},
	// 	AllowMethods:    []string{"GET", "POST", "PUT", "PATCH"},
	// }))

	engine.POST("/create-streamed", streamedHandler.CreateStreamed)
	engine.GET("/list-streamed", streamedHandler.GetListStreamed)
	engine.GET("/get-streamed/:id", streamedHandler.GetDetailStreamed)
	engine.PUT("/update-streamed/:id", streamedHandler.UpdateStreamed)
	engine.DELETE("/delete-streamed/:id", streamedHandler.DeleteStreamed)

	return engine
}

// func Router() *mux.Router {

// 	router := mux.NewRouter()
// 	//api tvseries_info
// 	router.HandleFunc("/api/tv", StreamedHandler.CreateStreamed).Methods("GET", "OPTIONS")
// 	router.HandleFunc("/api/tv/{id}", controller.GetTV).Methods("GET", "OPTIONS")
// 	router.HandleFunc("/api/tv", controller.NewTV).Methods("POST", "OPTIONS")
// 	router.HandleFunc("/api/tv/{id}", controller.UpdateTVNew).Methods("PUT", "OPTIONS")
// 	router.HandleFunc("/api/tv/{id}", controller.RemoveTV).Methods("DELETE", "OPTIONS")
// 	//api detailed
// 	router.HandleFunc("/api/detail", controller.GetDetailedAll).Methods("GET", "OPTIONS")
// 	router.HandleFunc("/api/detail", controller.NewDetailed).Methods("POST", "OPTIONS")
// 	router.HandleFunc("/api/detail/{id}", controller.RemoveDetailed).Methods("DELETE", "OPTIONS")

// 	return router
// }
