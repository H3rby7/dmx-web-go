package setup

import (
	"errors"
	"net/http"

	apiv1 "github.com/H3rby7/dmx-web-go/internal/api/v1"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

/*
Configure and start the webserver

  - Configuration is done with respect to the app options
*/
func SetUpAndStartServer() *http.Server {
	opts := options.GetAppOptions()

	router := gin.Default()
	if opts.Static != "" {
		router.Static("", opts.Static)
	}
	apiv1.RegisterHandlers(router.Group("/api/v1"))
	addr := ":" + opts.HttpPort
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	go func() {
		log.Warnf("Serving at %s", addr)
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()
	return srv
}
