package yeelight

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
)

type RawResponse struct {
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  json.RawMessage `json:"error"`
}

func (rr RawResponse) String() string {
	panic("implement me")
}

func (rr RawResponse) ToError() error {
	panic("implement me")
}

func (rr RawResponse) IsOk() bool {
	return bytes.Equal(rr.Result, json.RawMessage(`["ok"]`))
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

	c.log("Raw() params:", string(b))

	r, err := c.transport(ctx, host, string(b))
	if err != nil {
		return RawResponse{}, err
	}

	c.log("Raw() result", string(r))

	var target RawResponse
	if err := json.Unmarshal(r, &target); err != nil {
		return RawResponse{}, err
	}

	return target, nil
}

func defaultTransport(ctx context.Context, host string, raw string) ([]byte, error) {
	const crlf = "\r\n"

	var d net.Dialer

	conn, err := d.DialContext(ctx, "tcp", host)
	if nil != err {
		return nil, fmt.Errorf("cannot open connection to %s. %s", host, err)
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
