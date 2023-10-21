package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	dmxconn "github.com/H3rby7/dmx-web-go/internal/dmx"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/dmx-web-go/internal/setup"

	log "github.com/sirupsen/logrus"
)

func main() {
	options.InitAppOptions()
	setup.SetUpLogging()
	setup.SetUpDMX()
	srv := setup.SetUpAndStartServer()

	handleShutdown(srv)
}

/*
Waits until receiving a shutdown command, then runs cleanup/shutdown calls
*/
func handleShutdown(srv *http.Server) {
	// Wait for quit signal(s)
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
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

	reader := dmxconn.GetReader()
	if reader != nil {
		log.Debugf("Shutting down DMX reader...")
		if err := reader.Disconnect(); err != nil {
			log.Fatal("Error disconnecting DMX reader:", err)
		} else {
			log.Infof("DMX reader was shut down gracefully")
		}
	}

	log.Debugf("Shutting down DMX writer...")
	w := dmxconn.GetWriter()
	shouldClear := options.GetAppOptions().DmxClearOnQuit
	if shouldClear {
		log.Infof("Clearing DMX output to zeros")
		w.ClearStage()
		w.Commit()
	} else {
		log.Debugf("Skipping DMX output cleanup")
	}
	if err := w.Disconnect(); err != nil {
		log.Fatal("Error disconnecting DMX writer:", err)
	} else {
		log.Infof("DMX writer was shut down gracefully")
	}
	log.Infof("Finished cleaning up")
}
