package lib

import (
	"bytes"

	"github.com/pkg/errors"
)

// Compression format expansion is necessary because programs are
// free to avoid using pointers in the message they generate but
// all of the consumers are required to understand arriving messages
// that contain pointers.
//
//     first to bits set to 1: pointer
//     first to bits set to 0: label
//     01 and 10 are reserved for future implementations.
//
type CompressedName struct {
	Offset    uint16
	IsPointer bool
}

func (c CompressedName) Marshal() (res []byte, err error) {
	var (
		buf             = new(bytes.Buffer)
		msg_0      byte = 0
		msg_1      byte = 0
		offsetHigh byte = 0
		offsetLow  byte = 0
	)

	if c.IsPointer {
		msg_0 = (3 << 6)
	}

	offsetHigh = uint8(c.Offset & uint16(masks[7]))
	offsetLow = uint8(c.Offset >> 8)

	msg_1 = offsetLow
	msg_0 |= offsetHigh

	buf.Write([]byte{
		msg_0,
		msg_1,
	})

	res = buf.Bytes()

	return
}

func (c CompressedName) ExpandCompressedName(msg []byte) (name string, err error) {
	// TODO
	return
}

// UnmarshalCompressedName takes care of populating a CompressedName
// struct with information gathered from a compressed domain name
// buffer.
//
//		msg[0]			msg[1]
//	_____________________ ||______________________
//      0  1  2  3  4  5  6  7  0  1  2  3  4  5  6  7
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    | 1  1|                OFFSET                   |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//
func UnmarshalCompressedName(msg []byte, c *CompressedName) (n int, err error) {
	if c == nil {
		err = errors.Errorf("CompressedName must be non-nil")
		return
	}

	if len(msg) != 2 {
		err = errors.Errorf(
			"unexpected message size %d - should be 2 bytes long",
			len(msg))
		return
	}

	var (
		pointerFlag byte  = 0
		highValue   uint8 = 0
		lowValue    uint8 = 0
	)

	pointerFlag = (msg[0] >> 6)
	if pointerFlag == 3 {
		c.IsPointer = true
	}

	highValue = (msg[0] & masks[5])
	lowValue = msg[1]

	c.Offset = uint16(highValue) | uint16(lowValue)

	n = 2

	return
}
