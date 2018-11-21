package main

import (
	"net/http"

	"github.com/makki0205/config"
	"github.com/makki0205/log"
	"github.com/pkg/errors"
	"github.com/tarm/serial"
)

var mainUrl = config.Env("main_url")
var SerialPort = config.Env("serial_port")

func main() {
	err := SendIsWearing(true)
	if err != nil {
		log.Err(err)
		panic(err)
	}
	c := &serial.Config{Name: SerialPort, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Err(err)
		panic(err)
	}
	var buf = make([]byte, 5)
	for {
		_, err := s.Read(buf)
		if err != nil {
			log.Err(err)
		}
		if buf[0] == byte(0) {
			SendIsWearing(false)
		} else {
			SendIsWearing(true)
		}
	}
}

func SendIsWearing(flg bool) error {
	flgS := "false"

	if flg {
		flgS = "true"
	}
	_, err := http.Get(mainUrl + "/is_wearing/" + flgS)
	if err != nil {
		return errors.New("hat Driver [mainが見つかりません]")
	}
	return nil
}
