package lib

// RR represents a resource record.
//
//                                   1  1  1  1  1  1
//     0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
//   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//   |                      NAME                     |
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
	TTL uint32

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

func UnmarshalRR(msg []byte, r *RR) (n int, err error) {
	//	msg[0] and msg[1] --> compressed shit

	r.TYPE = QType(uint16(msg[3]) | uint16(msg[2]))
	r.CLASS = QClass(uint16(msg[5]) | uint16(msg[4]))
	r.TTL = uint32(msg[9]) | uint32(msg[8]) | uint32(msg[7]) | uint32(6)
	r.RDLENGTH = uint16(msg[11]) | uint16(msg[10])
	r.RDATA = msg[12 : 12+r.RDLENGTH]

	n = int(12 + r.RDLENGTH)
	return
}

func (r *RR) Marshal() (res []byte, err error) {
	return
}
