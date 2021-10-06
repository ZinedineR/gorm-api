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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api-gorm-setting/configuration/config"
	"api-gorm-setting/configuration/handler/http"
	"api-gorm-setting/configuration/repository"
	"api-gorm-setting/service"
)

func main() {
	log.Info().Msg("api-gorm-setting starting")
	cfg, err := config.NewConfig(".env")
	checkError(err)

	// tool.ErrorClient = setupErrorReporting(context.Background(), cfg)

	var db *gorm.DB
	db = openDatabase(cfg)

	defer func() {
		if sqlDB, err := db.DB(); err != nil {
			log.Fatal().Err(err)
			panic(err)
		} else {
			_ = sqlDB.Close()
		}
	}()
	// TVHandler := buildTVHandler(db)
	// engineTV := http.NewGinEngineTV(TVHandler, cfg.InternalConfig.Username, cfg.InternalConfig.Password)
	// serverTV := &nethttp.Server{
	// 	Addr:    fmt.Sprintf(":%s", cfg.Port),
	// 	Handler: engineTV,
	// }
	// // setGinMode(cfg.Env)
	// runServer(serverTV)
	// waitForShutdown(serverTV)

	StreamedHandler := buildStreamedHandler(db)
	engineStreamed := http.NewGinEngineStreamed(StreamedHandler, cfg.InternalConfig.Username, cfg.InternalConfig.Password)
	serverStreamed := &nethttp.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: engineStreamed,
	}
	// setGinMode(cfg.Env)
	runServer(serverStreamed)
	waitForShutdown(serverStreamed)
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
	log.Info().Msg("shutting down codelabs-service")

	// The context is used to inform the server it has 2 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("codelabs-service forced to shutdown")
	}

	log.Info().Msg("codelabs-service exiting")
}

func openDatabase(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	checkError(err)
	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func buildTVHandler(db *gorm.DB) *http.TVHandler {
	repo := repository.NewTVRepository(db)
	TVService := service.NewTVService(repo)
	return http.NewTVHandler(TVService)
}

func buildStreamedHandler(db *gorm.DB) *http.StreamedHandler {
	repo := repository.NewStreamedRepository(db)
	StreamedService := service.NewStreamedService(repo)
	return http.NewStreamedHandler(StreamedService)
}
