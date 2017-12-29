package lib

import (
	"github.com/pkg/errors"
)

// Message represents a DNS message that is sent
// via a transport.
type Message struct {

	// Header is always present. It includes fields
	// that specify which of the remaining sections
	// are present, and also specifies whether the
	// message is a query or a response, a standard
	// query of some other opcode etc
	Header

	// Question carries the "question" of the query,
	// defining parameters that determines what's
	// being asked.
	Questions []*Question

	// Answer
	// Authority
	// Additional
}

func (m Message) Marshal() (res []byte, err error) {
	var (
		questionPayload []byte
	)

	res, err = m.Header.Marshal()
	if err != nil {
		err = errors.Wrapf(err,
			"failed to create header payload %+v",
			m.Header)
		return
	}

	for _, question := range m.Questions {
		questionPayload, err = question.Marshal()
		if err != nil {
			err = errors.Wrapf(err,
				"failed to marshal question %+v",
				question)
			return
		}

		res = append(res, questionPayload...)
	}

	return
}
