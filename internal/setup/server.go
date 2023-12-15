// Package setup contains configurations for logging and server as well as service creation
package setup

import (
	"errors"
	"net/http"

	apiv1 "github.com/H3rby7/dmx-web-go/internal/api/v1"
	models_services "github.com/H3rby7/dmx-web-go/internal/model/services"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Configure and start the webserver in its own go-routine.
//
// Configuration is done with respect to the [AppOptions].
func SetUpAndStartServer(services *models_services.ApplicationServices) *http.Server {
	opts := options.GetAppOptions()

	log.Debugf("Creating router... ")
	router := gin.Default()

	if opts.Static != "" {
		log.Debugf("Adding static route for '%s'", opts.Static)
		router.Static("", opts.Static)
	}

	apiv1.RegisterHandlers(router.Group("/api/v1"), services)

	log.Tracef("Configuring HTTP Server... ")
	addr := ":" + opts.HttpPort
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Debugf("Starting HTTP server in separate go routine... ")
	go func() {
		log.Warnf("Serving at %s", addr)
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Errorf("listen: %s\n", err)
		}
	}()
	return srv
}
