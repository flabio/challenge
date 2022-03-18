package routers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestShowAll(t *testing.T) {
	r := gin.New()

	r.GET("/", adminController.All)

}
