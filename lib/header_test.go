package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaderMarshallingAndUnmarshalling(t *testing.T) {
	var testCases = []struct {
		desc        string
		header      *Header
		shouldError bool
	}{
		{
			desc:   "0-ed case",
			header: &Header{},
		},
		{
			desc: "case all set",
			header: &Header{
				ID:      2,
				QR:      1,
				Opcode:  OpcodeQuery,
				AA:      1,
				TC:      1,
				RD:      1,
				RA:      1,
				Z:       1,
				RCODE:   1,
				QDCOUNT: 1,
				ANCOUNT: 1,
				NSCOUNT: 1,
				ARCOUNT: 1,
			},
		},
		{
			desc: "some set",
			header: &Header{
				ID:      333,
				QR:      1,
				Opcode:  OpcodeQuery,
				AA:      1,
				TC:      0,
				RD:      1,
				RA:      0,
				Z:       1,
				RCODE:   1,
				QDCOUNT: 0,
				ANCOUNT: 1,
				NSCOUNT: 0,
				ARCOUNT: 2,
			},
		},
	}

	var (
		msg                []byte
		err                error
		unarmshalledHeader *Header
	)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			msg, err = tc.header.Marshal()
			require.NoError(t, err)
			assert.Equal(t, 12, len(msg))

			unarmshalledHeader = new(Header)
			err = UnmarshalHeader(msg, unarmshalledHeader)
			require.NoError(t, err)

			assert.Equal(t, tc.header.ID, unarmshalledHeader.ID)
			assert.Equal(t, tc.header.QR, unarmshalledHeader.QR)
		})
	}
}
