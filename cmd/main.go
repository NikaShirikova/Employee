package main

import (
	"context"
	"employee/internal/database/postgresql"
	"employee/internal/handler"
	"employee/internal/service"
	"employee/server"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	LoggerZap, err := zap.NewProduction()
	if err != nil {
		return
	}

	defer LoggerZap.Sync()

	if err := initConfig(); err != nil {
		LoggerZap.Fatal("Error initializing config ", zap.String("error ", err.Error()))
	}

	cfg := postgresql.Config{
		Host:        viper.GetString("db.host"),
		Username:    viper.GetString("db.username"),
		Password:    viper.GetString("db.password"),
		DBName:      viper.GetString("db.dbname"),
		SSLMode:     viper.GetString("db.sslmode"),
		TablePrefix: viper.GetString("db.tableprefix"),
	}
	db, err := postgresql.Init(cfg)
	if err != nil {
		LoggerZap.Fatal("Connect or initialize to database fatal ", zap.String("error ", err.Error()))
	}

	repos := postgresql.NewRepository(db, LoggerZap)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, LoggerZap)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes(), LoggerZap); err != nil {
			LoggerZap.Fatal("Start server to fatal ", zap.String("error ", err.Error()))
		}
	}()

	<-done
	LoggerZap.Info("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		LoggerZap.Fatal("Server shutdown failed ", zap.String("error ", err.Error()))
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
