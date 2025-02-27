package api

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRouter(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {

	server := &http.Server{
		Addr:         ":8080",
		Handler:      a.router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
