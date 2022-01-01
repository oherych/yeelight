package yeelight

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type RawResponse struct {
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  json.RawMessage `json:"error"`
}

func (rr RawResponse) ToError() error {
	if rr.Error == nil {
		return nil
	}

	if bytes.Equal(rr.Error, json.RawMessage(`{"code":-1, "message":"method not supported"}`)) {
		return ErrMethodNotSupported
	}

	return UnknownError(rr.Error)
}

func (rr RawResponse) Bind(target interface{}) error {
	err := json.Unmarshal(rr.Result, target)
	if err != nil {
		return ErrResponseJsonSyntax
	}

	return nil
}

func (c Client) Raw(ctx context.Context, host string, id int, method string, params ...interface{}) (RawResponse, error) {
	if params == nil {
		params = []interface{}{}
	}

	payload := map[string]interface{}{"id": id, "method": method, "params": params}

	b, err := json.Marshal(payload)
	if err != nil {
		return RawResponse{}, err
	}

	r, err := c.transport(ctx, host, string(b))
	if err != nil {
		return RawResponse{}, err
	}
	var target RawResponse
	if err := json.Unmarshal(r, &target); err != nil {
		return RawResponse{}, ErrResponseJsonSyntax
	}

	return target, nil
}

func (c Client) rawWithOk(ctx context.Context, host string, requestID int, method string, params ...interface{}) error {
	d, err := c.Raw(ctx, host, requestID, method, params...)
	if err != nil {
		return err
	}

	return d.ToError()
}

func defaultTransport(ctx context.Context, host string, raw string) ([]byte, error) {
	const crlf = "\r\n"

	var d net.Dialer

	conn, err := d.DialContext(ctx, "tcp", host)
	if err != nil {
		return nil, processDialError(err)
	}

	if _, err := fmt.Fprint(conn, raw+crlf); err != nil {
		return nil, err
	}

	res, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("cannot read command result %s", err)
	}

	return res, nil
}

func processDialError(err error) error {
	e, ok := err.(*net.OpError)
	if !ok {
		// return as isRaw
		return err
	}

	if ae, ok := e.Err.(*net.AddrError); ok {
		if ae.Err == "missing port in address" {
			return ErrMissingPortInAddress
		}
	}

	if se, ok := e.Err.(*os.SyscallError); ok {
		if se.Syscall == "connect" {
			return ErrConnect
		}
	}

	// return as isRaw
	return err
}
