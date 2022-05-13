package http

import (
	"embed"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

// INFO static files serve will not work if in folder found `index.html`.
func addStaticRoutes(router *gin.RouterGroup, prefix string, staticFS *embed.FS) (int, error) {
	fd, err := staticFS.ReadDir(prefix)
	if err != nil {
		return 0, err
	}

	for _, v := range fd {
		name := v.Name()
		if v.IsDir() {
			router.GET("/"+name+"/*filepath", func(c *gin.Context) {
				c.FileFromFS(path.Join(prefix, name, c.Param("filepath")), http.FS(staticFS))
			})
		} else {
			router.GET("/"+name, func(c *gin.Context) {
				c.FileFromFS(path.Join(prefix, name), http.FS(staticFS))
			})
		}
	}

	return len(fd), nil
}
