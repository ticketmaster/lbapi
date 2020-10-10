package main

import (
	"crypto/tls"
	"flag"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	filter "github.com/ticketmaster/authentication/gin"
	"github.com/ticketmaster/lbapi/common"
	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/env"
	"github.com/ticketmaster/lbapi/golog"
	"github.com/ticketmaster/lbapi/handler"
	"github.com/ticketmaster/lbapi/routeconfig"
)

func main() {
	////////////////////////////////////////////////////////////////////////////
	var err error
	env.Set()
	flag.Parse()
	log := logrus.New()
	////////////////////////////////////////////////////////////////////////////
	err = common.SetSources()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Load Balancer Sources enumerated.")
	////////////////////////////////////////////////////////////////////////////
	options := filter.NewAuthenticationOptions()
	options.ConfigPath = "etc"
	////////////////////////////////////////////////////////////////////////////
	router := gin.New()

	corsConfig := cors.DefaultConfig()
	if len(config.GlobalConfig.Lbm.CorsAllowedOrigins) > 0 {
		corsConfig.AllowOrigins = config.GlobalConfig.Lbm.CorsAllowedOrigins
	} else {
		corsConfig.AllowOrigins = []string{"http://localhost"}
	}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type", "Cache-Control", "Pragma"}
	router.Use(cors.New(corsConfig))
	////////////////////////////////////////////////////////////////////////////
	err = filter.UseAuthentication(router, options)
	if err != nil {
		logrus.Fatal(err)
	}
	////////////////////////////////////////////////////////////////////////////
	router.Use(golog.Logger(log))
	router.HEAD("/", func(c *gin.Context) {})
	router.GET("/", func(c *gin.Context) {})
	////////////////////////////////////////////////////////////////////////////
	v1 := router.Group("/api/v1")
	////////////////////////////////////////////////////////////////////////////

	l := routeconfig.NewLoadBalancer()
	_, err = handler.New(l, v1)
	if err != nil {
		logrus.Fatal(err)
	}
	v := routeconfig.NewVirtualServer()
	_, err = handler.New(v, v1)
	if err != nil {
		logrus.Fatal(err)
	}
	r := routeconfig.NewRecycle()
	_, err = handler.New(r, v1)
	if err != nil {
		logrus.Fatal(err)
	}
	////////////////////////////////////////////////////////////////////////////
	if config.GlobalConfig.Lbm.RunTLS {
		server := http.Server{
			Addr:         ":8443",
			Handler:      router,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
			TLSConfig: &tls.Config{
				MinVersion:               tls.VersionTLS12,
				CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
				PreferServerCipherSuites: true,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
					tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				},
			},
		}
		logrus.Fatal(server.ListenAndServeTLS("etc/"+config.GlobalConfig.Lbm.PemFile, "etc/"+config.GlobalConfig.Lbm.KeyFile))
	} else {
		server := http.Server{
			Addr:    ":8080",
			Handler: router,
		}
		logrus.Fatal(server.ListenAndServe())
	}
}
