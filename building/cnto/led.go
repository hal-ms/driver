package cnto

import (
	"net/http"
	"strconv"
	"time"

	"github.com/makki0205/config"
	"github.com/tarm/serial"

	"github.com/gin-gonic/gin"
)

var LEDA, _ = NewLED(config.Env("led_port_a"))
var LEDB, _ = NewLED(config.Env("led_port_b"))
var LEDC, _ = NewLED(config.Env("led_port_c"))
var LEDD, _ = NewLED(config.Env("led_port_d"))

func Led(c *gin.Context) {
	scene, err := strconv.Atoi(c.Param("scene"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "すうじのみ")
	}

	SendAll(uint8(scene))
	c.JSON(http.StatusOK, "OK")
}

func SendAll(scene uint8) {
	LEDA.Send(scene)
	LEDB.Send(scene)
	LEDC.Send(scene)
	LEDD.Send(scene)
}

type LED struct {
	p *serial.Port
}

func NewLED(port string) (*LED, error) {

	c := &serial.Config{Name: port, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	time.Sleep(2 * time.Second)
	return &LED{p: s}, nil
}

func (l *LED) Send(scene uint8) error {
	var res []byte
	res = append(res, ToByte(scene))
	_, err := l.p.Write(res)
	if err != nil {
		return err
	}
	err = l.p.Flush()
	return err
}

func ToByte(i uint8) byte {
	return byte(i)
}
