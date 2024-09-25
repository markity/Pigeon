package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"pigeon/im-gateway/protocol"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3501")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			_, err := conn.Write(protocol.PackData(protocol.MustEncodePacket(&protocol.HeartbeatPacket{})))
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second)
		}
	}()

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
				panic("check me")
			}
			fmt.Printf("recv %T, %v\n", pack, pack)
			switch packet := pack.(type) {
			case *protocol.S2CLoginRespPacket:
				fmt.Printf("\tcode:%v sessionId:%v version:%v sessions:%v\n", packet.Code, packet.SessionId, packet.Version, packet.Sessions)
				for _, v := range packet.Sessions {
					fmt.Printf("\t\tsessionId:%v loginAt:%v DeviceDesc:%v\n", v.SessionId, v.LoginAt, v.DeviceDesc)
				}
			case *protocol.S2CDeviceInfoBroadcastPacket:
				fmt.Printf("\tversion:%v devices:%v\n", packet.Version, packet.Devices)
				for _, v := range packet.Devices {
					fmt.Printf("\t\tsessionId:%v loginAt:%v DeviceDesc:%v\n", v.SessionId, v.LoginAt, v.DeviceDesc)
				}
			}
		}
	}()

	loginPacket := &protocol.C2SLoginPacket{
		Username:   "markity",
		Password:   "mark2004",
		DeviceDesc: "Android 5.14",
	}
	loginPacket.SetEchoCode("114514")
	_, err = conn.Write(protocol.PackData(protocol.MustEncodePacket(loginPacket)))
	if err != nil {
		panic(err)
	}
	select {}
}
