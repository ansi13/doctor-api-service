package ping

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	log.Println("It is working!!")
	c.IndentedJSON(http.StatusOK, map[string]string{"message": "pong"})
}
