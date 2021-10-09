package main

import (
	"context"
	"fmt"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"gorm.io/gorm"

	"api-gorm-setting/configuration/config"
	"api-gorm-setting/configuration/handler/http"
	"api-gorm-setting/configuration/repository"
	"api-gorm-setting/service"
)

func main() {
	log.Info().Msg("api-gorm-setting starting")
	cfg, err := config.NewConfig(".env")
	config.CheckError(err)

	// tool.ErrorClient = setupErrorReporting(context.Background(), cfg)

	var db *gorm.DB
	db = config.OpenDatabase(cfg)

	defer func() {
		if sqlDB, err := db.DB(); err != nil {
			log.Fatal().Err(err)
			panic(err)
		} else {
			_ = sqlDB.Close()
		}
	}()
	TVHandler := buildTVHandler(db)
	StreamedHandler := buildStreamedHandler(db)
	WatchedHandler := buildWatchedHandler(db)
	DetailedHandler := buildDetailedHandler(db)
	ActorHandler := buildActorHandler(db)
	LoginHandler := &http.Loginhandler{}
	UserHandler := buildUserHandler(db)
	engine := http.NewGinEngine(TVHandler, StreamedHandler, WatchedHandler, DetailedHandler, ActorHandler, LoginHandler, UserHandler, cfg.InternalConfig.Username, cfg.InternalConfig.Password)
	server := &nethttp.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: engine,
	}
	// setGinMode(cfg.Env)
	// authenticate()
	runServer(server)
	waitForShutdown(server)
}

func runServer(srv *nethttp.Server) {
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != nethttp.ErrServerClosed {
			log.Fatal().Err(err)
		}
	}()
}

func waitForShutdown(server *nethttp.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("shutting down api-gorm-setting")

	// The context is used to inform the server it has 2 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("api-gorm-setting forced to shutdown")
	}

	log.Info().Msg("api-gorm-setting exiting")
}

// func openDatabase(config *config.Config) *gorm.DB {
// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		config.Database.Host,
// 		config.Database.Port,
// 		config.Database.Username,
// 		config.Database.Password,
// 		config.Database.Name)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	checkError(err)
// 	return db
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
func buildTVHandler(db *gorm.DB) *http.TVHandler {
	repo := repository.NewTVRepository(db)
	TVService := service.NewTVService(repo)
	return http.NewTVHandler(TVService)
}

func buildDetailedHandler(db *gorm.DB) *http.DetailedHandler {
	repo := repository.NewDetailedRepository(db)
	DetailedService := service.NewDetailedService(repo)
	return http.NewDetailedHandler(DetailedService)
}

func buildActorHandler(db *gorm.DB) *http.ActorHandler {
	repo := repository.NewActorRepository(db)
	ActorService := service.NewActorService(repo)
	return http.NewActorHandler(ActorService)
}

func buildStreamedHandler(db *gorm.DB) *http.StreamedHandler {
	repo := repository.NewStreamedRepository(db)
	StreamedService := service.NewStreamedService(repo)
	return http.NewStreamedHandler(StreamedService)
}

func buildWatchedHandler(db *gorm.DB) *http.WatchedHandler {
	repo := repository.NewWatchedRepository(db)
	WatchedService := service.NewWatchedService(repo)
	return http.NewWatchedHandler(WatchedService)
}
func buildUserHandler(db *gorm.DB) *http.UserHandler {
	repo := repository.NewUserRepository(db)
	UserService := service.NewUserService(repo)
	return http.NewUserHandler(UserService)
}
