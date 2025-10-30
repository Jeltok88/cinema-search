package server

import (
	"github.com/gin-gonic/gin"
)

func NewGin() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/v1/films", func(c *gin.Context) { c.JSON(200, []any{}) })
	r.GET("/v1/playlist", func(c *gin.Context) { c.JSON(200, []any{}) })
	r.GET("/v1/actor", func(c *gin.Context) { c.JSON(200, []any{}) })
	r.GET("/v1/director", func(c *gin.Context) { c.JSON(200, []any{}) })
	r.GET("/health", func(c *gin.Context) { c.String(200, "ok") })
	return r
}
