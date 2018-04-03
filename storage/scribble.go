package storage

import (
	"github.com/nanobox-io/golang-scribble"
	"github.com/gin-gonic/gin"
)

func SetDBtoContext(db *scribble.Driver) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}

func DBInstance(c *gin.Context) *scribble.Driver {
	return c.MustGet("DB").(*scribble.Driver)
}