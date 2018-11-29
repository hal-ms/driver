package cnto

import (
	"net/http"

	"github.com/makki0205/config"

	"github.com/gin-gonic/gin"
	led "github.com/makki0205/led"
)

type LedInfo struct {
	Red   byte `json:"r"`
	Green byte `json:"g"`
	Blue  byte `json:"b"`
}

var LEDA, _ = led.NewLED(config.Env("led_port_a"))
var LEDB, _ = led.NewLED(config.Env("led_port_b"))
var LEDC, _ = led.NewLED(config.Env("led_port_c"))
var LEDD, _ = led.NewLED(config.Env("led_port_d"))

func Led(c *gin.Context) {
	var l LedInfo

	err := c.BindJSON(&l)
	if err != nil {
		c.JSON(http.StatusBadRequest, "NG")
		return
	}

	SendAll(l.Red, l.Green, l.Blue)
	c.JSON(http.StatusOK, "OK")
}

func SendAll(r, g, b uint8) {
	LEDA.Send(r, g, b)
	LEDB.Send(r, g, b)
	LEDC.Send(r, g, b)
	LEDD.Send(r, g, b)
}
