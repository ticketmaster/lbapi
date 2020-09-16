package main

import (
	"flag"

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
	////////////////////////////////////////////////////////////////////////////
	err = common.SetSources()
	if err != nil {
		logrus.Fatal(err)
	}
	////////////////////////////////////////////////////////////////////////////
	log := logrus.New()
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
		router.RunTLS(":8443", "etc/"+config.GlobalConfig.Lbm.PemFile, "etc/"+config.GlobalConfig.Lbm.KeyFile)
	} else {
		router.Run(":8080")
	}
}
