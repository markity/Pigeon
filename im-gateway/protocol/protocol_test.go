package protocol

import (
	"testing"
)

func TestProtocol(t *testing.T) {
	MustEncodePacket(&C2SLoginPacket{
		Username: "markity",
		Password: "password",
	})
}
