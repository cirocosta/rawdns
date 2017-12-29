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
	// QNAME refers to the domain name to be resolved in the query.
	// It's represented as a sequence of labels where each label
	// consists
	QNAME string

	// QTYPE specifies the type of the query to perform.
	QTYPE uint16

	// QCLASS
	QCLASS uint16
}

type QType uint16

const (
	QTypeUnknown QType = iota

	// Host address
	QTypeA

	// Authoritative name server
	QTypeNS

	QTypeMD
	QTypeMF

	// Canonical name for an alias
	QTypeCNAME

	// Marks the start of a zone of authority
	QTypeSOA

	QTypeMB
	QTypeMG
	QTypeMR
	QTypeNULL
	QTypeWKS

	// Domain name pointer
	QTypePTR
	QTypeHINFO
	QTypeMINFO

	// Mail exchange
	QTypeMX
	QTypeTXT
	QTypeAXFR  QType = 252
	QTypeMAILB QType = 253
	QTypeMAILA QType = 254

	// All records
	QTypeWildcard QType = 255
)

type QClass uint16

const (
	QClassUnknown QClass = iota

	// Internet
	QClassIN

	QClassCS
	QClassCH
	QClassHS

	// Any class
	QClassWildcard QClass = 255
)

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

func UnmarshalQuestion(msg []byte, q *Question) (err error) {
	if q == nil {
		err = errors.Errorf("question must be non-nil")
		return
	}

	return
}
