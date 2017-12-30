package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRRMarshallingAndUnmarshalling(t *testing.T) {
	var testCases = []struct {
		desc       string
		entity     *RR
		shouldFail bool
	}{
		{
			desc: "0-ed case",
			entity: &RR{
				NAME:     "test.com",
				RDLENGTH: 4,
				RDATA:    []byte{192, 168, 0, 1},
			},
		},
	}

	var (
		msg          []byte
		err          error
		unmarshalled *RR
	)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			msg, err = tc.entity.Marshal()

			require.NoError(t, err)

			unmarshalled = new(RR)
			_, err = UnmarshalRR(msg, unmarshalled)
			require.NoError(t, err)

			assert.Equal(t, tc.entity.RDATA, unmarshalled.RDATA)
			assert.Equal(t, tc.entity.TTL, unmarshalled.TTL)
			assert.Equal(t, tc.entity.RDLENGTH, unmarshalled.RDLENGTH)
			assert.Equal(t, tc.entity.TYPE, unmarshalled.TYPE)

		})
	}
}
