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

// CallResponse contains raw response from device.
type CallResponse struct {
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  json.RawMessage `json:"error"`
}

// ToError checks if the answer contains an error.
func (rr CallResponse) ToError() error {
	if rr.Error == nil {
		return nil
	}

	if bytes.Equal(rr.Error, json.RawMessage(`{"code":-1, "message":"method not supported"}`)) {
		return ErrMethodNotSupported
	}

	return UnknownError(rr.Error)
}

// Bind response result to provided variable
func (rr CallResponse) Bind(target interface{}) error {
	err := json.Unmarshal(rr.Result, target)
	if err != nil {
		return ErrResponseJSONSyntax
	}

	return nil
}

// Call method of device with provided parameters
func (c Client) Call(ctx context.Context, method string, params ...interface{}) (CallResponse, error) {
	if params == nil {
		params = []interface{}{}
	}

	// TODO: remove hardcoded id
	payload := map[string]interface{}{"id": 123, "method": method, "params": params}

	b, err := json.Marshal(payload)
	if err != nil {
		return CallResponse{}, err
	}

	r, err := c.transport(ctx, c.host, string(b))
	if err != nil {
		return CallResponse{}, err
	}
	var target CallResponse
	if err := json.Unmarshal(r, &target); err != nil {
		return CallResponse{}, ErrResponseJSONSyntax
	}

	return target, nil
}

func (c Client) rawWithOk(ctx context.Context, method string, params ...interface{}) error {
	d, err := c.Call(ctx, method, params...)
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

	defer conn.Close()

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
