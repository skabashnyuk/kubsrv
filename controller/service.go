package controller

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
	"os"
)

func GetService(c *gin.Context) {
	name := c.Param("name")
	name = strings.Replace(name,".",string(os.PathSeparator), -1)
	version := c.Param("version")
	c.File(filepath.Join("/Users/sj/dev/src/skabashnyuk/che-registry/", name, version, "CheService.yaml"))
}
