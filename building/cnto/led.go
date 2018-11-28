package cnto

import (
	"net/http"

	"github.com/gin-gonic/gin"
	led "github.com/makki0205/led"
)

type LedInfo struct {
	Red   byte `json:"r"`
	Green byte `json:"g"`
	Blue  byte `json:"b"`
}

var LED, _ = led.NewLED("COM9")

func Led(c *gin.Context) {
	var l LedInfo

	err := c.BindJSON(&l)
	if err != nil {
		c.JSON(http.StatusBadRequest, "NG")
		return
	}

	LED.Send(l.Red, l.Green, l.Blue)
	c.JSON(http.StatusOK, "OK")
}
