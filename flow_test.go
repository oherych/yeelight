package yeelight

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_StartColorFlow(t *testing.T) {
	var testCount = 12
	var testAction = FlowActionTurnOff
	var testExpressions = []FlowExpression{
		{
			Duration:   time.Minute,
			Mode:       FlowModeColor,
			Value:      30,
			Brightness: 40,
		},
		{
			Duration:   time.Second,
			Mode:       FlowColorTemperature,
			Value:      30,
			Brightness: 40,
		},
	}

	tests := map[string]struct {
		expressions []FlowExpression
		tr          transportFn

		expErr error
	}{
		"correct": {
			expressions: testExpressions,
			tr:          isRaw(t, testResultOkStr, `{"id":123,"method":"start_cf","params":[12,2,"60000,1,30,40,1000,2,30,40"]}`),
		},
		"empty_expressions": {
			expressions: []FlowExpression{},
			expErr:      ErrRequiredMinimumOneExpression,
		},
		"err_connection": {
			expressions: testExpressions,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.StartColorFlow(testCtx, testCount, testAction, tt.expressions)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_StartBackgroundColorFlow(t *testing.T) {
	var testCount = 12
	var testAction = FlowActionTurnOff
	var testExpressions = []FlowExpression{
		{
			Duration:   time.Minute,
			Mode:       FlowModeColor,
			Value:      30,
			Brightness: 40,
		},
		{
			Duration:   time.Second,
			Mode:       FlowColorTemperature,
			Value:      30,
			Brightness: 40,
		},
	}

	tests := map[string]struct {
		expressions []FlowExpression
		tr          transportFn

		expErr error
	}{
		"correct": {
			expressions: testExpressions,
			tr:          isRaw(t, testResultOkStr, `{"id":123,"method":"bg_start_cf","params":[12,2,"60000,1,30,40,1000,2,30,40"]}`),
		},
		"empty_expressions": {
			expressions: []FlowExpression{},
			expErr:      ErrRequiredMinimumOneExpression,
		},
		"err_connection": {
			expressions: testExpressions,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.StartBackgroundColorFlow(testCtx, testCount, testAction, tt.expressions)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_StopColorFlow(t *testing.T) {
	tests := map[string]struct {
		tr transportFn

		expErr error
	}{
		"correct": {
			tr:     isRaw(t, testResultOkStr, `{"id":123,"method":"stop_cf","params":[]}`),
			expErr: nil,
		},
		"err_connection": {
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.StopColorFlow(testCtx)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_StopBackgroundColorFlow(t *testing.T) {
	tests := map[string]struct {
		tr transportFn

		expErr error
	}{
		"correct": {
			tr:     isRaw(t, testResultOkStr, `{"id":123,"method":"bg_stop_cf","params":[]}`),
			expErr: nil,
		},
		"err_connection": {
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.StopBackgroundColorFlow(testCtx)

			require.Equal(t, tt.expErr, err)
		})
	}
}
