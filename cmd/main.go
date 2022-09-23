package main

import (
	"Employee/internal/database/postgresql"
	"Employee/internal/handler"
	"Employee/internal/service"
	"Employee/server"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := initConfig(); err != nil {
		panic("Error initializing config")
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
		panic("Connect to database fatal")
	}

	repos := postgresql.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			panic(err.Error())
		}
	}()

	<-done
	fmt.Println("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		panic("Server shutdown failed")
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
