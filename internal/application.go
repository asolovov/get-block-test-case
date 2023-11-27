package internal

import (
	"os"
	"os/signal"
	"syscall"

	"get-block-test-case/config"
	"get-block-test-case/pkg/logger"
)

// App is an application implementation
type App struct {
	config *config.Scheme
}

// NewApplication is used to get new App instance
func NewApplication() (*App, error) {
	a := &App{
		config: &config.Scheme{},
	}

	return a, nil
}

// Stop is used to gracefully stop the application
func (a *App) Stop() {
	logger.Log().WithField("layer", "App").Info("application stop...")
}

// Init is used to initialize the application and all dependencies
func (a *App) Init() error {
	logger.Log().WithField("layer", "App").Info("application init...")
	return nil
}

// Serve is used to run the application
func (a *App) Serve() error {

	// Gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-quit

	a.Stop()
	return nil
}

// Config is used to get application configurations
func (a *App) Config() *config.Scheme {
	return a.config
}
