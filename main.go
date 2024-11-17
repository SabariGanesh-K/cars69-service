package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"database/sql"
 _ "github.com/golang/mock/mockgen/model"
"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// "github.com/hibiken/asynq"
	// "github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/SabariGanesh-K/cars69-service.git/api"
	db "github.com/SabariGanesh-K/cars69-service.git/db/sqlc"
	// "github.com/SabariGanesh-K/cars69-service.git/email"
	"github.com/SabariGanesh-K/cars69-service.git/util"
	// "github.com/SabariGanesh-K/cars69-service.git/redis-worker"
	"golang.org/x/sync/errgroup"

)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	fmt.Printf("hello")
	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	conn,err:= sql.Open(config.DBDriver,config.DBSource)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	// runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)

	


	waitGroup, ctx := errgroup.WithContext(ctx)

	runGinServer(config,store)


	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}


func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}
	// corsConfig := cors.Config{
	// 	AllowOrigins:  []string{"*"},
	// 	AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
	// 	AllowHeaders:  []string{"Content-Type", "Authorization", "Accept"},
	// 	ExposeHeaders: []string{"Content-Length"},
	// 	MaxAge:        12 * time.Hour,
	// }
	// server.Router.Use(cors.New(corsConfig))


	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
