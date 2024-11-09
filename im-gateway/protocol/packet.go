package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/markity/go-reactor/pkg/buffer"
)

// 包 = 4字节包头 + 包体, 包头用小端

// 包的长度有软限制, 4m
const MaxDataLength = 4194304

// 封装一个包, 理论上length为0的包也能算包
func PackData(data []byte) []byte {
	buf := bytes.Buffer{}
	bs := binary.LittleEndian.AppendUint32(nil, uint32(len(data)))
	buf.Write(bs)
	buf.Write(data)
	return buf.Bytes()
}

// 从buffer中解析出一个包
// 4字节包头 + 包体, 函数返回值为包体字节
// 如果包的长度太长, error != nil
func UnpackDataFromBuffer(buf buffer.Buffer) ([]byte, bool, error) {
	if buf.ReadableBytes() <= 4 {
		return nil, false, nil
	}

	bs := buf.Peek()[:4]
	l := binary.LittleEndian.Uint32(bs)
	if l > uint32(MaxDataLength) {
		return nil, false, errors.New("packet too big")
	}
	if buf.ReadableBytes() < int(4+l) {
		return nil, false, nil
	}

	data := buf.Peek()[4 : l+4]

	buf.Retrieve(int(4 + l))
	return data, true, nil
}
