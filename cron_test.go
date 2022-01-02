package yeelight

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_AddCron(t *testing.T) {
	testTimeout := (30 * time.Minute) + (5 * time.Second)

	tests := map[string]struct {
		on bool

		tr transportFn

		expErr error
	}{
		"correct_on": {
			on: true,
			tr: isRaw(t, testResultOkStr, `{"id":123,"method":"cron_add","params":[1,30]}`),
		},
		"correct_off": {
			on: false,
			tr: isRaw(t, testResultOkStr, `{"id":123,"method":"cron_add","params":[0,30]}`),
		},
		"err_connection": {
			on: true,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.AddCron(testCtx, tt.on, testTimeout)

			require.Equal(t, tt.expErr, err)
		})
	}
}
func TestClient_GetCron(t *testing.T) {
	tests := map[string]struct {
		on bool

		tr transportFn

		exp    time.Duration
		expErr error
	}{
		"correct_on": {
			on:  true,
			tr:  isRaw(t, `{"id":1, "result":[{"type": 1, "delay": 15, "mix": 0}]}`, `{"id":123,"method":"cron_get","params":[1]}`),
			exp: 15 * time.Minute,
		},
		"correct_off": {
			on:  false,
			tr:  isRaw(t, `{"id":1, "result":[]}`, `{"id":123,"method":"cron_get","params":[0]}`),
			exp: 0,
		},
		"err_connection": {
			on: true,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
		"json_error": {
			on: true,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return []byte(`{"id":1, "result":[{"delay" : "im_bad_value"}]}`), nil
			},
			expErr: ErrResponseJsonSyntax,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			got, err := Client{host: testHost, transport: tt.tr}.GetCron(testCtx, tt.on)

			require.Equal(t, tt.expErr, err)
			require.Equal(t, tt.exp, got)
		})
	}
}

func TestClient_DeleteCron(t *testing.T) {
	tests := map[string]struct {
		on bool

		tr transportFn

		expErr error
	}{
		"correct_on": {
			on: true,
			tr: isRaw(t, testResultOkStr, `{"id":123,"method":"cron_del","params":[1]}`),
		},
		"correct_off": {
			on: false,
			tr: isRaw(t, testResultOkStr, `{"id":123,"method":"cron_del","params":[0]}`),
		},
		"err_connection": {
			on: true,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.DeleteCron(testCtx, tt.on)

			require.Equal(t, tt.expErr, err)
		})
	}
}
