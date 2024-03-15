package main

import (
	"context"
	"curr-quote/internal/adapters/exchange"
	"curr-quote/internal/app"
	"curr-quote/internal/ports/httpserver"
	"curr-quote/internal/repo"
	"curr-quote/pkg/logger"
	"errors"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SetConfigs(configPath string) error {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("cannot read config file %w", err)
	}
	return nil
}

func ConnectToPostgres(ctx context.Context) (*pgx.Conn, error) {
	adsRepoUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("postgres-quotes-db.username"),
		viper.GetString("postgres-quotes-db.password"),
		viper.GetString("postgres-quotes-db.host"),
		viper.GetInt("postgres-quotes-db.port"),
		viper.GetString("postgres-quotes-db.dbname"),
		viper.GetString("postgres-quotes-db.sslmode"))

	// 30 attempts to connect to postgres when starting in docker container
	for i := 0; i < 30; i++ {
		conn, err := pgx.Connect(ctx, adsRepoUrl)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			return conn, nil
		}
	}

	return nil, errors.New("unable to connect to postgres ads repo")
}

const (
	dockerConfigFile = "config/config-docker.yml"
	localConfigFile  = "config/config-local.yml"
)

func main() {
	ctx := context.Background()
	logs := logger.New()

	isDocker := flag.Bool("docker", false, "flag if this project is running in docker container")
	flag.Parse()
	var configPath string
	if *isDocker {
		configPath = dockerConfigFile
	} else {
		configPath = localConfigFile
	}

	if err := SetConfigs(configPath); err != nil {
		logs.Fatal(nil, fmt.Sprintf("reading configs: %s", err.Error()))
	}

	conn, err := ConnectToPostgres(ctx)
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("connecting to postgres: %s", err.Error()))
	}
	defer func() {
		_ = conn.Close(ctx)
	}()
	logs.Info(nil, "successfully connected to postgres")

	srvAddr := fmt.Sprintf("%s:%d",
		viper.GetString("http-server.host"),
		viper.GetInt("http-server.port"),
	)

	r := repo.New(conn)
	a := app.New(ctx, exchange.New(), r, logs)

	srv := httpserver.New(srvAddr, a, logs)

	go func() {
		_ = srv.ListenAndServe()
	}()
	logs.Info(nil, "http server successfully started")
	// preparing graceful shutdown
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT)

	// waiting for Ctrl+C
	<-osSignals

	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second) // 30s timeout to finish all active connections
	defer cancel()

	_ = srv.Shutdown(shutdownCtx)
	logs.Info(nil, "successfully stopped http server")
}
