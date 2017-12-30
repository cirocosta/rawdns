package lib

import (
	"bytes"
	"encoding/binary"

	"github.com/pkg/errors"
)

// RCODE denotes a 4bit field that specifies the response
// code for a query.
type RCODE byte

const (
	RCODENoError RCODE = iota
	RCODEFormatError
	RCODEServerFailure
	RCODENameError
	RCODENotImplemented
	RCODERefused
)

// Opcode denotes a 4bit field that specified the query type.
type Opcode byte

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
	QR byte

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
	AA byte

	// TC indicates whether the message was (T)run(C)ated due to the length
	// being grater than the permitted on the transmission channel.
	// 1bit.
	TC byte

	// RD indicates whether (R)ecursion is (D)esired or not.
	// 1bit.
	RD byte

	// RA indidicates whether (R)ecursion is (A)vailable or not.
	// 1bit.
	RA byte

	// Z is reserved for future use
	Z byte

	// RCODE contains the (R)esponse (CODE) - it's a 4bit field that is
	// set as part of responses.
	RCODE byte

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

func UnmarshalHeader(msg []byte, h *Header) (n int, err error) {
	if h == nil {
		err = errors.Errorf("header must be non-nil")
		return
	}

	if len(msg) != 12 {
		err = errors.Errorf(
			"msg does not have the expected size - %d",
			len(msg))
		return
	}

	var (
		h1_0 byte = 0
		h1_1 byte = 0
	)

	// ID is formed by the first two bytes such that
	// it results in an uint16.
	// As this comes from the network in UDP packets we can assume that it comes
	// in BigEndian (network byte order), thus, consider the first byte of each
	// the most significant.
	h.ID = uint16(msg[1]) | uint16(msg[0])<<8

	// take the first byte of the second row (QR, Opcode, AA, TC and RD)
	h1_0 = msg[2]

	// Each value is got from right bitshifting
	// followed by masking such that we end up
	// with only that number.
	// Here we're starting from right to left.
	h.RD = h1_0 & masks[0]
	h.TC = (h1_0 >> 1) & masks[0]
	h.AA = (h1_0 >> 2) & masks[0]
	h.Opcode = Opcode((h1_0 >> 3) & masks[3])
	h.QR = (h1_0 >> 7) & masks[0]

	// take the second byte of the second row (RA, Z, RCODE)
	h1_1 = msg[3]
	h.RCODE = h1_1 & masks[3]
	h.Z = (h1_1 >> 4) & masks[2]
	h.RA = (h1_1 >> 7) & masks[0]

	// QDCOUNT, ANCOUNT, NSCOUNT and ARCOUNT are all formed by two bytes that
	// results in uint16.
	// As this comes from the network in UDP packets we can assume that it comes
	// in BigEndian (network byte order), thus, consider the first byte of each
	// the most significant.
	h.QDCOUNT = uint16(msg[5]) | uint16((msg[4] << 8))

	// ANCOUNT is formed by two bytes that results in uint16
	h.ANCOUNT = uint16(msg[7]) | uint16(msg[6]<<8)

	// NSCOUNT is formed by two bytes that results in uint16
	h.NSCOUNT = uint16(msg[9]) | uint16(msg[8]<<8)

	// ARCOUNT is formed by two bytes that results in uint16
	h.ARCOUNT = uint16(msg[11]) | uint16(msg[10]<<8)

	n = 12
	return
}

func (h Header) Marshal() (res []byte, err error) {
	var (
		buf       = new(bytes.Buffer)
		h1_0 byte = 0
		h1_1 byte = 0
	)

	binary.Write(buf, binary.BigEndian, h.ID)

	// first 8bit part of the second row
	// QR :		0
	// Opcode:	1 2 3 4
	// AA:		5
	// TC:		6
	// RD:		7
	h1_0 = h.QR << (7 - 0)
	h1_0 |= byte(h.Opcode) << (7 - (1 + 3))
	h1_0 |= h.AA << (7 - 5)
	h1_0 |= h.TC << (7 - 6)
	h1_0 |= h.RD << (7 - 7)

	// second 8bit part of the second row
	// RA:		0
	// Z:		1 2 3
	// RCODE:	4 5 6 7
	h1_1 = h.RA << (7 - 0)
	h1_1 |= h.Z << (7 - 1)
	h1_1 |= byte(h.RCODE) << (7 - (4 + 3))

	buf.WriteByte(h1_0)
	buf.WriteByte(h1_1)

	binary.Write(buf, binary.BigEndian, h.QDCOUNT)
	binary.Write(buf, binary.BigEndian, h.ANCOUNT)
	binary.Write(buf, binary.BigEndian, h.NSCOUNT)
	binary.Write(buf, binary.BigEndian, h.ARCOUNT)

	res = buf.Bytes()
	return
}
