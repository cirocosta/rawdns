package lib

import ()

//
//
//                                   1  1  1  1  1  1
//     0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
//   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//   |                                               |
//   /                                               /
//   /                      NAME                     /
//   |                                               |
//   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//   |                      TYPE                     |
//   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//   |                     CLASS                     |
//   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//   |                      TTL                      |
//   |                                               |
//   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//   |                   RDLENGTH                    |
//   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--|
//   /                     RDATA                     /
//   /                                               /
//   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//
type RR struct {
	// Domain name to which this resource record pertains
	NAME string

	// The type of the RR. It specifies the meaning of the
	// data that RDATA contains.
	TYPE QType

	// Identifies the class of the data holded on RDATA.
	CLASS QClass

	// Indicates the time interval in seconds that the resource
	// recorded may be cached before it should be discarded.
	TTL uint16

	// Specifies the length in octets of the RDATA field.
	// ps.: if there's a pointer in the RDATA, this length
	// will not count the final result (expanded), but the
	// actual amount in transfer.
	RDLENGTH uint16

	// Generic data from the record.
	// The format of the information contained here varies
	// according to the tupple {TYPE, CLASS} of the RR.
	RDATA []byte
}

//
// TODO - Compression format expansion
//
// Necessary because programs are free to avoid using pointers
// in the message they generate but all of the consumers are
// required to understand arriving messages that contain pointers.
//
// first to bits set to 1: pointer
// first to bits set to 0: label
//
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    | 1  1|                OFFSET                   |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//
//
//

type RRA struct {
	RR
}

func (r RRA) ParseRDATA() (err error) {
	return
}
