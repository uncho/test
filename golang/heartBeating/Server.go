/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  17/09/2017 14:57
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */
package main

import (
	"github.com/astaxie/beego"
	"net"
	"os"
	"time"
)

func main() {
	server := ":8090"
	netListen, err := net.Listen("tcp", server)
	if err != nil {
		beego.Warning("Fatal error : ", err)
		os.Exit(1)
	}
	beego.Info("waiting for client")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			beego.Warning("Fatal error : ", err)
			return
		}
		conn.SetReadDeadline(time.Now().Add(time.Duration(8) * time.Second))
		beego.Trace(conn.RemoteAddr().String(), "-->>connect success")
		go HandleConnection(conn)
	}

}

func HandleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			beego.Warning("Fatal error : ", err)
			return
		}
		data := buffer[:n]
		message := make(chan byte)
		go GetMessage(data, message)
		go HeartBeat(conn, message, 4)
		beego.Trace(conn.RemoteAddr().String(), string(buffer[:n]))
	}
	defer conn.Close()
}

func GetMessage(bytes []byte, message chan byte) {
	for _, v := range bytes {
		message <- v
	}
	close(message)
}

func HeartBeat(conn net.Conn, message chan byte, timeout int) {
	select {
	case <-message:
		//beego.Trace(conn.RemoteAddr().String(), string(fk), ": ", "Heart beating ")
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	case <-time.After(time.Second * 5):
		beego.Warning(conn.RemoteAddr().String(), "time out")
		conn.Close()
	}
}
