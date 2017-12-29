package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompressedNameMarshallingAndUnmarshalling(t *testing.T) {
	var testCases = []struct {
		desc   string
		entity *CompressedName
	}{
		{
			desc: "pointer",
			entity: &CompressedName{
				IsPointer: true,
				Offset:    10,
			},
		},
	}

	var (
		msg          []byte
		err          error
		unmarshalled *CompressedName
	)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			msg, err = tc.entity.Marshal()
			require.NoError(t, err)

			unmarshalled = new(CompressedName)
			err = UnmarshalCompressedName(msg, unmarshalled)
			require.NoError(t, err)

			assert.Equal(t, tc.entity.IsPointer, unmarshalled.IsPointer)
			assert.Equal(t, tc.entity.Offset, unmarshalled.Offset)
		})
	}

}
