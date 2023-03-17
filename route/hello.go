package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "Hello, HomeChores!")
}
