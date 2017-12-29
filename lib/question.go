package lib

import (
	"bytes"
	"encoding/binary"
	"strings"

	"github.com/pkg/errors"
)

//                                    1  1  1  1  1  1
//      0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                                               |
//    /                     QNAME                     /
//    /                                               /
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                     QTYPE                     |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                     QCLASS                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//
type Question struct {
	QNAME  string
	QTYPE  uint16
	QCLASS uint16
}

func (q Question) Marshal() (res []byte, err error) {
	var (
		buf    = new(bytes.Buffer)
		labels []string
	)

	// split QNAME by label using `.` as the separator

	labels = strings.Split(q.QNAME, ".")
	if len(labels) < 2 {
		err = errors.Errorf(
			"malformed qname %s",
			q.QNAME)
		return
	}

	for _, label := range labels {
		buf.WriteByte(uint8(len(label)))
		buf.Write([]byte(label))
	}

	buf.WriteByte(0)

	binary.Write(buf, binary.BigEndian, q.QTYPE)
	binary.Write(buf, binary.BigEndian, q.QNAME)

	res = buf.Bytes()
	return

}
