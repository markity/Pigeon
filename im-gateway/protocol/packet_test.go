package protocol

import (
	"testing"

	"github.com/markity/go-reactor/pkg/buffer"
)

func TestPacket(t *testing.T) {
	buf := buffer.NewBuffer()
	buf.Append(PackData([]byte{1, 2, 3, 4}))
	buf.Append(PackData([]byte{5, 6, 7, 8}))

	bs, ok, err := UnpackDataFromBuffer(buf)
	if string(bs) != string([]byte{1, 2, 3, 4}) || !ok || err != nil {
		t.Errorf("fail %v", bs)
	}
	bs, ok, err = UnpackDataFromBuffer(buf)
	if string(bs) != string([]byte{5, 6, 7, 8}) || !ok || err != nil {
		t.Errorf("fail %v", bs)
	}
}
