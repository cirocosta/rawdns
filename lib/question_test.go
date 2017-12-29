package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuestionMarshallingAndUnmarshalling(t *testing.T) {
	var testCases = []struct {
		desc   string
		entity *Question
	}{
		{
			desc:   "0-ed case",
			entity: &Question{},
		},
	}

	var (
		msg          []byte
		err          error
		unmarshalled *Question
	)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			msg, err = tc.entity.Marshal()
			require.NoError(t, err)
			assert.Equal(t, 12, len(msg))

			unmarshalled = new(Question)
			err = UnmarshalQuestion(msg, unmarshalled)
			require.NoError(t, err)
		})
	}
}
