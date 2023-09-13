package gq

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"net"
)

// Client for accessing kdb/q. Use [NewClient] or [NewClientTLS] to connect to a q server.
type Client struct {
	conn     net.Conn
	capacity Capacity
}

// handleshake see [NewClient] for details.
func (client *Client) handleshake(ctx context.Context, auth string, capacity Capacity) error {
	buf := bytes.NewBufferString(auth)
	buf.WriteByte(byte(capacity))
	buf.WriteByte(0)

	_, err := client.conn.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to send authentication info: %w", err)
	}

	reply := make([]byte, 2+len(auth))

	n, err := client.conn.Read(reply)
	if err != nil {
		return fmt.Errorf("failed to read authentication reply: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("failed to authenticate. max version: %d", reply[0])
	}

	client.capacity = Capacity(reply[0])

	return nil
}

// NewClient creates a new client without tls support, auth can be empty.
//
// The procedure to connect to the server is described [on kx.com].
// The handshake is achieved by sending "username:password" plus the capacity byte, and the null-terminator.
// (Basically a null-terminated c string with the last char being the capacity.
// If authentication fails, the connection will be closed right away. If successful,
// a single byte of reply is sent by indicating the capacity of the server.
//
// For capacity, please use [Default_Capacity] = 3 if the [Capacity] of the server is unknown.
//
// [on kx.com]: https://code.kx.com/q/basics/ipc/
func NewClient(ctx context.Context, addr string, auth string, capacity Capacity) (*Client, error) {
	dialer := net.Dialer{}

	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s due to: %w", addr, err)
	}

	client := &Client{conn: conn}
	err = client.handleshake(ctx, auth, capacity)
	if err != nil {
		defer conn.Close()
		return nil, err
	}

	return client, nil
}

// NewClientTLS creates a new client with TLS support, auth or config can be nil or empty.
//
// For capacity, please use [DefaultCapacity] = 3 if the [Capacity] of the server is unknown.
//
// See [NewClient] for details.
func NewClientTLS(ctx context.Context, addr string, auth string, capacity Capacity, config *tls.Config) (*Client, error) {
	dialer := tls.Dialer{Config: config}
	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s due to: %w", addr, err)
	}

	client := &Client{conn: conn}
	err = client.handleshake(ctx, auth, capacity)
	if err != nil {
		defer conn.Close()
		return nil, err
	}

	return client, nil
}

// CloseClient close connection of the client.
func (client *Client) Close() error {
	if client.conn != nil {
		err := client.conn.Close()
		client.conn = nil
		if err != nil {
			return err
		}
	}

	return nil
}
