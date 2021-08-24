//+build !js

package websocket

import (
	"context"
	"fmt"
	"io"

	"github.com/fasthttp/websocket"
)

func NewClient(ctx context.Context, url string) (*Client, error) {
	d := websocket.Dialer{}
	conn, _, err := d.DialContext(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("dial %q: %w", url, err)
	}

	c := &Client{
		conn: conn,
	}

	return c, nil
}

type Client struct {
	conn  *websocket.Conn
	codec Codec
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Write(data interface{}) (err error) {
	var wc io.WriteCloser
	wc, err = c.conn.NextWriter(websocket.BinaryMessage)
	if err != nil {
		return fmt.Errorf("next message writer: %w", err)
	}

	defer func() {
		closeErr := wc.Close()
		if err == nil && closeErr != nil {
			err = fmt.Errorf("close message writer: %w", err)
		}
	}()

	err = c.codec.NewEncoder(wc).Encode(data)
	if err != nil {
		fmt.Errorf("encode %s message: %w", c.codec.Name(), err)
	}

	return
}

func (c *Client) Read(data interface{}) (err error) {
	var r io.Reader
	_, r, err = c.conn.NextReader()
	if err != nil {
		return fmt.Errorf("next read message: %w", err)
	}

	err = c.codec.NewDecoder(r).Decode(data)
	if err != nil {
		return fmt.Errorf("decode %s message: %w", c.codec.Name(), err)
	}

	return
}
