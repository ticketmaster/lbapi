package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/ticketmaster/lbapi/userenv"
	"github.com/gin-gonic/gin"
)

// Create ...
func (h Handler) Create(c *gin.Context) {
	////////////////////////////////////////////////////////////////////////////
	oUser := userenv.New(c)
	////////////////////////////////////////////////////////////////////////////
	p, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Error(err)
	}
	////////////////////////////////////////////////////////////////////////////
	handler, ok := h.Definition.(Create)
	if !ok {
		c.Error(errors.New("the handler definition does not contain a Post method"))
	}
	////////////////////////////////////////////////////////////////////////////
	r, err := handler.Create(p, oUser)
	if err != nil {
		c.Error(err)
		c.Status(400)
	}
	////////////////////////////////////////////////////////////////////////////
	if err := json.NewEncoder(c.Writer).Encode(r); err != nil {
		c.Error(err)
	}
}

// ImportAll ...
func (h Handler) ImportAll(c *gin.Context) {
	////////////////////////////////////////////////////////////////////////////
	oUser := userenv.New(c)
	////////////////////////////////////////////////////////////////////////////
	handler, ok := h.Definition.(ImportAll)
	if !ok {
		c.Error(errors.New("the handler definition does not contain a ImportAll method"))
	}
	////////////////////////////////////////////////////////////////////////////
	r, err := handler.ImportAll(oUser)
	if err != nil {
		c.Status(400)
		c.Error(err)
	}
	////////////////////////////////////////////////////////////////////////////
	if err := json.NewEncoder(c.Writer).Encode(r); err != nil {
		c.Error(err)
	}
}

// Delete ...
func (h Handler) Delete(c *gin.Context) {
	////////////////////////////////////////////////////////////////////////////
	oUser := userenv.New(c)
	////////////////////////////////////////////////////////////////////////////
	filter := c.Param("id")
	////////////////////////////////////////////////////////////////////////////
	handler, ok := h.Definition.(Delete)
	if !ok {
		c.Error(errors.New("the handler definition does not contain a Delete method"))
	}
	////////////////////////////////////////////////////////////////////////////
	r, err := handler.Delete(filter, oUser)
	if err != nil {
		c.Status(400)
		c.Error(err)
	}
	////////////////////////////////////////////////////////////////////////////
	if err := json.NewEncoder(c.Writer).Encode(r); err != nil {
		c.Error(err)
	}
}

// Fetch ...
func (h Handler) Fetch(c *gin.Context) {
	////////////////////////////////////////////////////////////////////////////
	oUser := userenv.New(c)
	////////////////////////////////////////////////////////////////////////////
	p := c.Request.URL.Query()
	////////////////////////////////////////////////////////////////////////////
	var toEncoder interface{}
	handler, ok := h.Definition.(Fetch)
	if !ok {
		c.Error(errors.New("the handler definition does not contain a Get method"))
	}
	r, err := handler.Fetch(p, 0, oUser)
	if err != nil {
		c.Status(400)
		c.Error(err)
	}
	toEncoder = r
	////////////////////////////////////////////////////////////////////////////
	if err := json.NewEncoder(c.Writer).Encode(toEncoder); err != nil {
		c.Error(err)
	}
}

// FetchVs ...
func (h Handler) FetchVs(c *gin.Context) {
	////////////////////////////////////////////////////////////////////////////
	oUser := userenv.New(c)
	////////////////////////////////////////////////////////////////////////////
	p := c.Request.URL.Query()
	////////////////////////////////////////////////////////////////////////////
	var toEncoder interface{}
	handler, ok := h.Definition.(FetchVirtualServices)
	if !ok {
		c.Error(errors.New("the handler definition does not contain a Get method"))
	}
	r, err := handler.FetchVirtualServices(p, 0, oUser)
	if err != nil {
		c.Status(400)
		c.Error(err)
	}
	toEncoder = r
	////////////////////////////////////////////////////////////////////////////
	if err := json.NewEncoder(c.Writer).Encode(toEncoder); err != nil {
		c.Error(err)
	}
}

// Backup ...
func (h Handler) Backup(c *gin.Context) {
	////////////////////////////////////////////////////////////////////////////
	oUser := userenv.New(c)
	////////////////////////////////////////////////////////////////////////////
	handler, ok := h.Definition.(Backup)
	if !ok {
		c.Error(errors.New("the handler definition does not contain a Get method"))
	}
	////////////////////////////////////////////////////////////////////////////
	err := handler.Backup(oUser)
	msg := "backup complete"
	if err != nil {
		c.Status(400)
		c.Error(err)
		msg = err.Error()
	}
	////////////////////////////////////////////////////////////////////////////
	if err := json.NewEncoder(c.Writer).Encode(msg); err != nil {
		c.Error(err)
	}
}

// FetchByID ...
func (h Handler) FetchByID(c *gin.Context) {
	////////////////////////////////////////////////////////////////////////////
	oUser := userenv.New(c)
	////////////////////////////////////////////////////////////////////////////
	filter := c.Param("id")
	handler, ok := h.Definition.(FetchByID)
	////////////////////////////////////////////////////////////////////////////
	if !ok {
		c.Error(errors.New("the handler definition does not contain a GetByID method"))
	}
	////////////////////////////////////////////////////////////////////////////
	r, err := handler.FetchByID(filter, oUser)
	if err != nil {
		c.Status(400)
		c.Error(err)
	}
	////////////////////////////////////////////////////////////////////////////
	if err := json.NewEncoder(c.Writer).Encode(r); err != nil {
		c.Error(err)
	}
}

// FetchStaged ...
func (h Handler) FetchStaged(c *gin.Context) {
	////////////////////////////////////////////////////////////////////////////
	oUser := userenv.New(c)
	////////////////////////////////////////////////////////////////////////////
	filter := c.Param("id")
	handler, ok := h.Definition.(FetchStaged)
	////////////////////////////////////////////////////////////////////////////
	if !ok {
		c.Error(errors.New("the handler definition does not contain a GetByID method"))
	}
	////////////////////////////////////////////////////////////////////////////
	r, err := handler.FetchStaged(filter, oUser)
	if err != nil {
		c.Status(200)
		c.Error(err)
	}
	////////////////////////////////////////////////////////////////////////////
	if err := json.NewEncoder(c.Writer).Encode(r); err != nil {
		c.Error(err)
	}
}

// Migrate ...
func (h Handler) Migrate(c *gin.Context) {
	////////////////////////////////////////////////////////////////////////////
	oUser := userenv.New(c)
	////////////////////////////////////////////////////////////////////////////
	filter := c.Param("id")
	handler, ok := h.Definition.(Migrate)
	////////////////////////////////////////////////////////////////////////////
	if !ok {
		c.Error(errors.New("the handler definition does not contain a GetByID method"))
	}
	////////////////////////////////////////////////////////////////////////////
	r, err := handler.Migrate(filter, oUser)
	if err != nil {
		c.Status(200)
		c.Error(err)
	}
	////////////////////////////////////////////////////////////////////////////
	if err := json.NewEncoder(c.Writer).Encode(r); err != nil {
		c.Error(err)
	}
}

// StageMigration ...
func (h Handler) StageMigration(c *gin.Context) {
	////////////////////////////////////////////////////////////////////////////
	oUser := userenv.New(c)
	////////////////////////////////////////////////////////////////////////////
	b, err := ioutil.ReadAll(c.Request.Body)
	filter := c.Param("id")
	handler, ok := h.Definition.(StageMigration)
	////////////////////////////////////////////////////////////////////////////
	if !ok {
		c.Error(errors.New("the handler definition does not contain a GetByID method"))
	}
	////////////////////////////////////////////////////////////////////////////
	r, err := handler.StageMigration(b, filter, oUser)
	if err != nil {
		c.Status(200)
		c.Error(err)
	}
	////////////////////////////////////////////////////////////////////////////
	if err := json.NewEncoder(c.Writer).Encode(r); err != nil {
		c.Error(err)
	}
}

// Modify ...
func (h Handler) Modify(c *gin.Context) {
	////////////////////////////////////////////////////////////////////////////
	oUser := userenv.New(c)
	////////////////////////////////////////////////////////////////////////////
	p, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Error(err)
	}
	////////////////////////////////////////////////////////////////////////////
	handler, ok := h.Definition.(Modify)
	if !ok {
		c.Error(errors.New("the handler definition does not contain a Put method"))
	}
	////////////////////////////////////////////////////////////////////////////
	r, err := handler.Modify(p, oUser)
	if err != nil {
		c.Status(400)
		c.Error(err)
	}
	////////////////////////////////////////////////////////////////////////////
	if err := json.NewEncoder(c.Writer).Encode(r); err != nil {
		c.Error(err)
	}
}

// New - package constructor.
func New(definition interface{}, route *gin.RouterGroup) (*Handler, error) {
	handler := Handler{Definition: definition}
	rh, ok := definition.(routeHandler)
	////////////////////////////////////////////////////////////////////////////
	if !ok {
		return nil, errors.New("specified handler does not implement the routeHandler interface")
	}
	////////////////////////////////////////////////////////////////////////////
	routeString := rh.GetRoute()
	if _, ok := definition.(Fetch); ok {
		route.GET("/"+routeString, handler.Fetch)
	}
	if _, ok := definition.(FetchByID); ok {
		route.GET("/"+routeString+"/:id", handler.FetchByID)
	}
	if _, ok := definition.(Modify); ok {
		route.PUT("/"+routeString, handler.Modify)
	}
	if _, ok := definition.(Create); ok {
		route.POST("/"+routeString, handler.Create)
	}
	if _, ok := definition.(Delete); ok {
		route.DELETE("/"+routeString+"/:id", handler.Delete)
	}
	if routeString != "recycle" {
		if _, ok := definition.(ImportAll); ok {
			route.POST("/source/"+routeString, handler.ImportAll)
		}
	}
	if routeString == "virtualserver" {
		if _, ok := definition.(Backup); ok {
			route.GET("/simple/"+routeString, handler.FetchVs)
		}
		if _, ok := definition.(Backup); ok {
			route.POST("/backup/"+routeString, handler.Backup)
		}
		if _, ok := definition.(StageMigration); ok {
			route.POST("/migrate/"+routeString+"/:id", handler.StageMigration)
		}
		if _, ok := definition.(Migrate); ok {
			route.PUT("/migrate/"+routeString+"/:id", handler.Migrate)
		}
		if _, ok := definition.(FetchStaged); ok {
			route.GET("/migrate/"+routeString+"/:id", handler.FetchStaged)
		}
	}
	return &handler, nil
}
