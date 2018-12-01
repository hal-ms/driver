package main

import (
	"fmt"
	"net/http"

	"github.com/makki0205/config"
	"github.com/makki0205/log"

	"github.com/tarm/serial"
)

var mainUrl = config.Env("main_url")
var SerialPort = config.Env("serial_port")

func main() {
	SendIsWearing(false)
	c := &serial.Config{Name: SerialPort, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Err(err)
		panic(err)
	}
	old := byte(0)
	for {
		var buf = make([]byte, 100)
		_, err := s.Read(buf)
		if err != nil {
			log.Err(err)
		}
		if old != buf[0] {
			//かぶっている
			SendIsWearing(buf[0] == byte(49))
			old = buf[0]
		}
		//	if buf[0] == byte(0) {
		//		if Cap.state == true {
		//			SendIsWearing(false)
		//		}
		//	} else {
		//		if Cap.state == false {
		//			SendIsWearing(true)
		//		}
		//	}
	}
}

func SendIsWearing(flg bool) {
	fmt.Println("send!!", flg)
	flgS := "false"
	if flg {
		flgS = "true"
	}

	go func() {
		_, err := http.Get(mainUrl + "/api/game/is_wearing/" + flgS)
		if err != nil {
			log.Err(err)
		}
	}()

}
