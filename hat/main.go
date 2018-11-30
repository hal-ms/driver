package main

import (
	"errors"
	"net/http"

	"github.com/makki0205/config"
	"github.com/makki0205/log"

	"github.com/tarm/serial"
)

var mainUrl = config.Env("main_url")
var SerialPort = config.Env("serial_port")

type cap struct {
	state bool
}

var Cap = cap{state: false}

func main() {
	err := SendIsWearing(false)
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
			if Cap.state == true {
				SendIsWearing(false)
			}
		} else {
			if Cap.state == false {
				SendIsWearing(true)
			}
		}
	}
}

func SendIsWearing(flg bool) error {
	Cap.state = flg
	flgS := "false"

	if flg {
		flgS = "true"
	}
	_, err := http.Get(mainUrl + "/api/game/is_wearing/" + flgS)
	if err != nil {
		return errors.New("hat Driver [mainが見つかりません]")
	}
	return nil
}
