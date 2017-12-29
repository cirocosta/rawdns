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

	// NAME is the Domain name to which this resource record pertains.
	// It might either come in the compressed format or not, depending
	// on the server.
	// Typically this should come compressed (indicated by the first two
	// bits).
	NAME string

	// TYPE is the type of the RR. It specifies the meaning of the
	// data that RDATA contains.
	TYPE QType

	// CLASS identifies the class of the data holded on RDATA.
	CLASS QClass

	// TTL indicates the time interval in seconds that the resource
	// recorded may be cached before it should be discarded.
	TTL uint16

	// RDLENGTH specifies the length in octets of the RDATA field.
	// ps.: if there's a pointer in the RDATA, this length
	// will not count the final result (expanded), but the
	// actual amount in transfer.
	RDLENGTH uint16

	// RDATA is the generic data from the record.
	// The format of the information contained here varies
	// according to the tupple {TYPE, CLASS} of the RR.
	RDATA []byte
}

func UnmarshalRR(msg []byte, r *RR) (err error) {
	if (msg[0] >> 7) > 0 {

	}

	return
}

func (r *RR) Marshal() (res []byte, err error) {
	return
}
