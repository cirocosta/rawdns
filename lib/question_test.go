package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuestionMarshallingAndUnmarshalling(t *testing.T) {
	var testCases = []struct {
		desc       string
		entity     *Question
		shouldFail bool
	}{
		{
			desc:       "0-ed case",
			entity:     &Question{},
			shouldFail: true,
		},
		{
			desc: "empty label should fail",
			entity: &Question{
				QNAME: "..",
			},
			shouldFail: true,
		},
		{
			desc: "well formed",
			entity: &Question{
				QNAME:  "test.com",
				QTYPE:  QTypeTXT,
				QCLASS: QClassIN,
			},
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
			if tc.shouldFail {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			unmarshalled = new(Question)
			err = UnmarshalQuestion(msg, unmarshalled)
			require.NoError(t, err)

			assert.Equal(t, tc.entity.QNAME, unmarshalled.QNAME)
			assert.Equal(t, tc.entity.QCLASS, unmarshalled.QCLASS)
			assert.Equal(t, tc.entity.QTYPE, unmarshalled.QTYPE)
		})
	}
}
