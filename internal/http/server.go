package http

import (
	"context"
	"embed"
	"net/http"
	"time"

	"github.com/archekb/lsx024b/internal/http/middleware"
	"github.com/archekb/lsx024b/internal/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Env struct {
	Release bool
	State   interface{}
}

// HTTPServer
type HTTPServer struct {
	router *gin.Engine
	server *http.Server
	env    *Env
}

// New make new HTTP Server instance, add routes, middlewares
func New(staticFS *embed.FS, env *Env) *HTTPServer {
	if env.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	hs := &HTTPServer{
		env:    env,
		router: router,
		server: &http.Server{
			Handler: router,
		},
	}

	// set middlewares
	hs.router.Use(middleware.Logger())
	hs.router.Use(gin.Recovery())
	hs.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		// AllowOriginFunc: func(origin string) bool {
		// 	return true
		// },
		MaxAge: 12 * time.Hour,
	}))

	// set no route handler
	hs.router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// set routes handlers
	hs.router.GET("/", func(c *gin.Context) {
		c.FileFromFS("/web/index.htm", http.FS(staticFS))
	})

	// add routes for embed content
	sr, err := addStaticRoutes(&hs.router.RouterGroup, "web", staticFS)
	if err != nil || sr < 1 {
		log.Warningln("No embed content was found")
	}

	api := hs.router.Group("/api")

	// controller api
	api.GET("/state", func(c *gin.Context) {
		c.JSON(http.StatusOK, hs.env.State)
	})

	return hs
}

// Run server
func (hs *HTTPServer) Run(address string) {
	hs.server.Addr = address
	go func() {
		if err := hs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP Server error: %s", err)
		}
	}()
}

// RunTLS server
func (hs *HTTPServer) RunTLS(address string, cert, key string) {
	hs.server.Addr = address
	go func() {
		if err := hs.server.ListenAndServeTLS(cert, key); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTPS Server error: %s", err)
		}
	}()
}

// Shutdown server
func (hs *HTTPServer) Shutdown() {

	// shutdown server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := hs.server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
