package lib

import (
	"net"

	"github.com/pkg/errors"
)

type Client struct {
	conn net.Conn
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

func (c *Client) LookupAddr(addr string) (ips []string, err error) {
	// create the query dns msg struct
	// create the answer dns msg struct
	// marshall the query
	// unmarshall the answer
	// grab the ips from the answer

	return
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}

	return
}
