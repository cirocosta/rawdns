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

	// Answers encapsulates the resource records
	// retrieved when receiving answers from the
	// server queried.
	Answers []*RR
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

func UnmarshalMessage(msg []byte, m *Message) (err error) {
	var (
		header    = &Header{}
		questions []*Question
		rrs       []*RR
		ndx       int = 0
		bytesRead int = 0
		n         int = 0
	)

	n, err = UnmarshalHeader(msg, header)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to read header")
		return
	}

	bytesRead += n

	questions = make([]*Question, header.QDCOUNT)
	for ndx, _ = range questions {
		questions[ndx] = new(Question)

		n, err = UnmarshalQuestion(msg[bytesRead:], questions[ndx])
		if err != nil {
			err = errors.Wrapf(err,
				"failed to read question %d",
				ndx)
			return
		}

		bytesRead += n
	}

	rrs = make([]*RR, header.ANCOUNT)
	for ndx, _ = range rrs {
		rrs[ndx] = new(RR)

		n, err = UnmarshalRR(msg[bytesRead:], rrs[ndx])
		if err != nil {
			err = errors.Wrapf(err,
				"failed to read answer %d",
				ndx)
			return
		}

		bytesRead += n
	}

	m.Header = *header
	m.Questions = questions
	m.Answers = rrs

	return
}
