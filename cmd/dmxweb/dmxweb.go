package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	models_services "github.com/H3rby7/dmx-web-go/internal/model/services"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/dmx-web-go/internal/setup"

	log "github.com/sirupsen/logrus"
)

func main() {
	options.InitAppOptions()
	setup.SetUpLogging()
	svcs := setup.InitServices()

	srv := setup.SetUpAndStartServer(svcs)

	handleShutdown(srv, svcs)
}

/*
Waits until receiving a shutdown command, then runs cleanup/shutdown calls
*/
func handleShutdown(srv *http.Server, services *models_services.ApplicationServices) {
	log.Tracef("Preparing to handle shutdown commands... ")
	// Wait for quit signal(s)
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Debugf("Ready to handle shutdown commands.")
	<-quit

	log.Infof("Received shutdown command - Starting to clean up...")

	grace := 5 * time.Second
	log.Debugf("Granting the server %v to shutdown", grace)
	ctx, cancel := context.WithTimeout(context.Background(), grace)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	} else {
		log.Infof("Server was shut down gracefully")
	}
	opts := options.GetAppOptions()
	if ok, _ := opts.CanWriteDMX(); ok {
		services.FadingService.Stop()
	}
	services.DMXReaderService.DisconnectDMX()
	services.FadingService.DisconnectDMX()

	log.Infof("Finished cleaning up")
}
