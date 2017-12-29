package lib

import (
	"bytes"
	"encoding/binary"

	"github.com/pkg/errors"
)

// RCODE denotes a 4bit field that specifies the response
// code for a query.
type RCODE uint8

const (
	RCODENoError RCODE = iota
	RCODEFormatError
	RCODEServerFailure
	RCODENameError
	RCODENotImplemented
	RCODERefused
)

// Opcode denotes a 4bit field that specified the query type.
type Opcode uint8

const (
	OpcodeQuery Opcode = iota
	OpcodeIquery
	OpcodeStatus
)

// Header encapsulates the construct of the header part of the DNS
// query message.
// It follows the conventions stated at RFC1035 section 4.1.1.
//
//
// The header contains the following fields:
//
//				0  1  2  3  4  5  6  7
//      0  1  2  3  4  5  6  7  8  9  A  B  C  D  E  F
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                      ID                       |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                    QDCOUNT                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                    ANCOUNT                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                    NSCOUNT                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                    ARCOUNT                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//
type Header struct {

	// ID is an arbitrary 16bit request identifier that is
	// forwarded back in the response so that we can match them up.
	ID uint16

	// QR is an 1bit flag specifying whether this message is a query (0)
	// of a response (1)
	// 1bit
	QR uint8

	// Opcode is a 4bit field that specifies the query type.
	// Possible values are:
	// 0		- standard query		(QUERY)
	// 1		- inverse query			(IQUERY)
	// 2		- server status request		(STATUS)
	// 3 to 15	- reserved for future use
	Opcode Opcode

	// AA indicates whether this is an (A)nswer from an (A)uthoritative
	// server.
	// Valid in responses only.
	// 1bit.
	AA uint8

	// TC indicates whether the message was (T)run(C)ated due to the length
	// being grater than the permitted on the transmission channel.
	// 1bit.
	TC uint8

	// RD indicates whether (R)ecursion is (D)esired or not.
	// 1bit.
	RD uint8

	// RA indidicates whether (R)ecursion is (A)vailable or not.
	// 1bit.
	RA uint8

	// Z is reserved for future use
	Z uint8

	// RCODE contains the (R)esponse (CODE) - it's a 4bit field that is
	// set as part of responses.
	RCODE uint8

	// QDCOUNT specifies the number of entries in the question section
	QDCOUNT uint16

	// ANCOUNT specifies the number of resource records (RR) in the answer
	// section
	ANCOUNT uint16

	// NSCOUNT specifies the number of name server resource records in the
	// authority section
	NSCOUNT uint16

	// ARCOUNT specifies the number of resource records in the additional
	// records section
	ARCOUNT uint16
}

func UnmarshalHeader(msg []byte, h *Header) (err error) {
	if h == nil {
		err = errors.Errorf("header must not be nil")
		return
	}

	return
}

func (h Header) Marshal() (res []byte, err error) {
	var (
		buf *bytes.Buffer

		// first 8bit part of the second row
		// QR :		0
		// Opcode:	1 2 3 4
		// AA:		5
		// TC:		6
		// RD:		7
		h1_0 uint8 = 0

		// second 8bit part of the second row
		// RA:		0
		// Z:		1 2 3
		// RCODE:	4 5 6 7
		h1_1 uint8 = 0
	)

	err = binary.Write(buf, binary.BigEndian, h.ID)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to write ID bytes %x into buffer",
			h.ID)
		return
	}

	h1_0 = h.QR << 7

	res = buf.Bytes()
	return
}
