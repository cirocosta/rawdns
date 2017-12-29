package lib

import (
	"net"
	"sync"

	"github.com/pkg/errors"
)

type Client struct {
	nextId uint16
	conn   net.Conn

	mu sync.Mutex
}

type ClientConfig struct {
	Address string
}

func NewClient(cfg ClientConfig) (c Client, err error) {
	if cfg.Address == "" {
		err = errors.Errorf("Address must be specified")
		return
	}

	c.conn, err = net.Dial("udp", cfg.Address)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to create connection to address %s",
			cfg.Address)
		return
	}

	return
}

// create the query dns msg struct
// create the answer dns msg struct
// marshall the query
// unmarshall the answer
// grab the ips from the answer
func (c *Client) LookupAddr(addr string) (ips []string, err error) {
	var (
		id      uint16
		payload []byte
	)

	func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		id = c.nextId
		c.nextId += 1
	}()

	queryMsg := &Message{
		Header: Header{
			ID: id,
		},
		Questions: []*Question{
			{
				QNAME:  "google.com",
				QTYPE:  QTypeA,
				QCLASS: QClassIN,
			},
		},
	}

	payload, _ = queryMsg.Marshal()
	_, err = c.conn.Write(payload)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to write query payload %+v",
			queryMsg)
		return
	}

	return
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}

	return
}
