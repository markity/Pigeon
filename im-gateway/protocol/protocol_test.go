package protocol

import (
	"testing"
)

func TestProtocol(t *testing.T) {
	p := &C2SLoginPacket{
		Username:   "markity",
		Password:   "password",
		DeviceDesc: "Android 5.1",
	}
	p.SetEchoCode("123123")

	b, err := ParseC2SPacket(MustEncodePacket(p))
	if err != nil {
		t.Fatalf("parse packet failed: " + err.Error())
	}

	bb := b.(*C2SLoginPacket)
	if bb.Username != p.Username {
		t.Fatalf("username not match: %s", bb.Username)
	}
	if bb.Password != p.Password {
		t.Fatalf("password not match: %s", bb.Password)
	}
	if bb.DeviceDesc != p.DeviceDesc {
		t.Fatalf("device desc not match: %s", bb.DeviceDesc)
	}
	if bb.EchoCode() != p.EchoCode() {
		t.Fatalf("echo code not match: %s", bb.EchoCode())
	}

	push := S2CPushMessagePacket{
		Data: map[string]interface{}{"test": "123", "hello": nil},
	}
	if string(MustEncodePacket(&push)) != `{"packet_type":"push-msg","data":{"hello":null,"test":"123"},"echo_code":""}` {
		t.Error("encode push packet failed")
	}
}
