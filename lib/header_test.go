package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaderMarshallingAndUnmarshalling(t *testing.T) {
	var testCases = []struct {
		desc   string
		entity *Header
	}{
		{
			desc:   "0-ed case",
			entity: &Header{},
		},
		{
			desc: "case all set",
			entity: &Header{
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
			entity: &Header{
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
		msg          []byte
		err          error
		unmarshalled *Header
	)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			msg, err = tc.entity.Marshal()
			require.NoError(t, err)
			assert.Equal(t, 12, len(msg))

			unmarshalled = new(Header)
			err = UnmarshalHeader(msg, unmarshalled)
			require.NoError(t, err)

			assert.Equal(t, tc.entity.ID, unmarshalled.ID)
			assert.Equal(t, tc.entity.QR, unmarshalled.QR)
			assert.Equal(t, tc.entity.Opcode, unmarshalled.Opcode)
			assert.Equal(t, tc.entity.AA, unmarshalled.AA)
			assert.Equal(t, tc.entity.RCODE, unmarshalled.RCODE)
			assert.Equal(t, tc.entity.QDCOUNT, unmarshalled.QDCOUNT)
			assert.Equal(t, tc.entity.ARCOUNT, unmarshalled.ARCOUNT)
		})
	}
}
