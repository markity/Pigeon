package main

import (
	"fmt"
	"net"
	"time"

	connkeeper "pigeon/client/conn_keeper"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3501")
	if err != nil {
		panic(err)
	}
	ck := connkeeper.NewConnKeeper(conn, connkeeper.Config{
		HbInterval: time.Second,
		HbTimeout:  time.Second * 3,
		Username:   "markity",
		Password:   "mark2004",
		DeviceDesc: "Android",
	})
	eventC := ck.Start()
	for {
		select {
		case event := <-eventC:
			fmt.Println(event)
		}
	}
}
