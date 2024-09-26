package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"pigeon/im-gateway/protocol"
	"sync"
	"time"

	interactive "github.com/markity/Interactive-Console"
)

func main() {
	promptStyle := interactive.GetDefaultSytleAttr()
	promptStyle.Bold = true
	promptStyle.Foreground = interactive.ColorDarkGreen
	win := interactive.Run(interactive.Config{
		Prompt:               '>',
		PromptStyle:          promptStyle,
		BlockInputAfterRun:   false,
		BlockInputAfterEnter: true,
		TraceAfterRun:        true,
		EventHandleMask:      interactive.EventMaskKeyCtrlC,
	})

	win.SendLineBack("connecting to server...")
	win.SetBlockInput(true)

	var conn net.Conn
	for {
		var err error
		conn, err = net.Dial("tcp", "127.0.0.1:3501")
		if err != nil {
			win.SendLineBack("fail, retrying: " + err.Error())
			time.Sleep(time.Second)
			continue
		}
		break
	}

	win.SendLineBack("connect ok!")

	uq := NewUnboundedQueen()
	cmdChan := win.GetCmdChan()
	eventChan := win.GetEventChan()
	heartbeatChan := time.Tick(time.Second)
	connErrChan := make(chan error, 2)
	packetRecvChan := make(chan interface{})
	// 负责发消息
	go func() {
		for {
			_, err := conn.Write(uq.PopBlock().([]byte))
			if err != nil {
				connErrChan <- err
				return
			}
		}
	}()
	// 负责收消息
	go func() {
		lengthData := [4]byte{}
		for {
			_, err := io.ReadFull(conn, lengthData[:])
			if err != nil {
				panic(err)
			}
			length := binary.LittleEndian.Uint32(lengthData[:])
			bs := make([]byte, length)
			_, err = io.ReadFull(conn, bs)
			if err != nil {
				panic(err)
			}

			pack, ok := protocol.ParseS2CPacket(bs)
			if !ok {
				connErrChan <- errors.New("parse packet error")
				conn.Close()
			}
			packetRecvChan <- pack
		}
	}()
	for {
		select {
		case cmd := <-cmdChan:
			win.SendLineBack(cmd)
			win.SetBlockInput(false)
		case ev := <-eventChan:
			switch ev.(type) {
			case *interactive.EventKeyCtrlC:
				win.Stop()
				return
			}
		case <-heartbeatChan:
			uq.Push(protocol.PackData(protocol.MustEncodePacket(&protocol.HeartbeatPacket{})))
		case err := <-connErrChan:
			win.SendLineBack("conn lost: " + err.Error())
			win.SetBlockInput(true)
		case packet := <-packetRecvChan:
			win.SendLineBack(fmt.Sprint(packet))
		}
	}
}

type unboundedQueen struct {
	queen []interface{}
	l     sync.Mutex
	cond  *sync.Cond
}

func NewUnboundedQueen() *unboundedQueen {
	uq := unboundedQueen{}
	uq.queen = make([]interface{}, 0)
	uq.cond = sync.NewCond(&uq.l)
	return &uq
}

func (uq *unboundedQueen) Push(i interface{}) {
	uq.l.Lock()
	uq.queen = append(uq.queen, i)
	uq.l.Unlock()
	uq.cond.Broadcast()
}

func (uq *unboundedQueen) PopBlock() interface{} {
	uq.l.Lock()
	for {
		if len(uq.queen) == 0 {
			uq.cond.Wait()
		} else {
			ret := uq.queen[0]
			uq.queen = uq.queen[1:]
			uq.l.Unlock()
			return ret
		}
	}
}
