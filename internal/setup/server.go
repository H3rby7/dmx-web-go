package setup

import (
	"errors"
	"net/http"

	apiv1 "github.com/H3rby7/dmx-web-go/internal/api/v1"
	"github.com/H3rby7/dmx-web-go/internal/options"
	"github.com/H3rby7/dmx-web-go/internal/services/chase"
	"github.com/H3rby7/dmx-web-go/internal/services/config"
	"github.com/H3rby7/dmx-web-go/internal/services/trigger"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func initServicesAndApi(router *gin.Engine) {
	configSvc := config.NewConfigService()
	chaseSvc := chase.NewChaseService(configSvc)
	triggerSvc := trigger.NewTriggerService(configSvc, chaseSvc)
	apiv1.RegisterHandlers(router.Group("/api/v1"), triggerSvc)
}

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
	initServicesAndApi(router)
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
