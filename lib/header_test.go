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
			desc: "ID gets properly splitted into the first two bytes",
			header: &Header{
				ID: 2,
			},
		},
	}

	var (
		msg []byte
		err error
	)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			msg, err = tc.header.Marshal()
			if tc.shouldError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, 12, len(msg))
		})
	}
}
