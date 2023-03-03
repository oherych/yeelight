package yeelight

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_SetAdjust(t *testing.T) {
	testAction := AdjustActionIncrease
	testProp := AdjustPropBright

	tests := map[string]struct {
		tr transportFn

		expErr error
	}{
		"correct": {
			tr: isRaw(t, testResultOkStr, `{"id":123,"method":"set_adjust","params":["increase","bright"]}`),
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
			err := Client{host: testHost, transport: tt.tr}.SetAdjust(testCtx, testAction, testProp)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetBackgroundAdjust(t *testing.T) {
	testAction := AdjustActionIncrease
	testProp := AdjustPropBright

	tests := map[string]struct {
		tr transportFn

		expErr error
	}{
		"correct": {
			tr: isRaw(t, testResultOkStr, `{"id":123,"method":"bg_set_adjust","params":["increase","bright"]}`),
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
			err := Client{host: testHost, transport: tt.tr}.SetBackgroundAdjust(testCtx, testAction, testProp)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_AdjustBright(t *testing.T) {
	tests := map[string]struct {
		percentage int
		duration   time.Duration

		tr transportFn

		expErr error
	}{
		"correct": {
			percentage: 50,
			duration:   time.Minute,
			tr:         isRaw(t, testResultOkStr, `{"id":123,"method":"adjust_bright","params":[50,60000]}`),
		},
		"wrong_percentage": {
			percentage: -150,
			duration:   time.Minute,
			expErr:     ErrWrongAdjustPercentage,
		},
		"wrong_duration": {
			percentage: 50,
			duration:   time.Millisecond,
			expErr:     ErrDurationTooSmall,
		},
		"err_connection": {
			percentage: 50,
			duration:   time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.AdjustBright(testCtx, tt.percentage, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_AdjustBackgroundBright(t *testing.T) {
	tests := map[string]struct {
		percentage int
		duration   time.Duration

		tr transportFn

		expErr error
	}{
		"correct": {
			percentage: 50,
			duration:   time.Minute,
			tr:         isRaw(t, testResultOkStr, `{"id":123,"method":"bg_adjust_bright","params":[50,60000]}`),
		},
		"wrong_percentage": {
			percentage: -150,
			duration:   time.Minute,
			expErr:     ErrWrongAdjustPercentage,
		},
		"wrong_duration": {
			percentage: 50,
			duration:   time.Millisecond,
			expErr:     ErrDurationTooSmall,
		},
		"err_connection": {
			percentage: 50,
			duration:   time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.AdjustBackgroundBright(testCtx, tt.percentage, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_AdjustColorTemperature(t *testing.T) {
	tests := map[string]struct {
		percentage int
		duration   time.Duration

		tr transportFn

		expErr error
	}{
		"correct": {
			percentage: 50,
			duration:   time.Minute,
			tr:         isRaw(t, testResultOkStr, `{"id":123,"method":"adjust_ct","params":[50,60000]}`),
		},
		"wrong_percentage": {
			percentage: -150,
			duration:   time.Minute,
			expErr:     ErrWrongAdjustPercentage,
		},
		"wrong_duration": {
			percentage: 50,
			duration:   time.Millisecond,
			expErr:     ErrDurationTooSmall,
		},
		"err_connection": {
			percentage: 50,
			duration:   time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.AdjustColorTemperature(testCtx, tt.percentage, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_AdjustBackgroundColorTemperature(t *testing.T) {
	tests := map[string]struct {
		percentage int
		duration   time.Duration

		tr transportFn

		expErr error
	}{
		"correct": {
			percentage: 50,
			duration:   time.Minute,
			tr:         isRaw(t, testResultOkStr, `{"id":123,"method":"bg_adjust_ct","params":[50,60000]}`),
		},
		"wrong_percentage": {
			percentage: -150,
			duration:   time.Minute,
			expErr:     ErrWrongAdjustPercentage,
		},
		"wrong_duration": {
			percentage: 50,
			duration:   time.Millisecond,
			expErr:     ErrDurationTooSmall,
		},
		"err_connection": {
			percentage: 50,
			duration:   time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.AdjustBackgroundColorTemperature(testCtx, tt.percentage, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_AdjustColor(t *testing.T) {
	tests := map[string]struct {
		percentage int
		duration   time.Duration

		tr transportFn

		expErr error
	}{
		"correct": {
			percentage: 50,
			duration:   time.Minute,
			tr:         isRaw(t, testResultOkStr, `{"id":123,"method":"adjust_color","params":[50,60000]}`),
		},
		"wrong_percentage": {
			percentage: -150,
			duration:   time.Minute,
			expErr:     ErrWrongAdjustPercentage,
		},
		"wrong_duration": {
			percentage: 50,
			duration:   time.Millisecond,
			expErr:     ErrDurationTooSmall,
		},
		"err_connection": {
			percentage: 50,
			duration:   time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.AdjustColor(testCtx, tt.percentage, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_AdjustBackgroundColor(t *testing.T) {
	tests := map[string]struct {
		percentage int
		duration   time.Duration

		tr transportFn

		expErr error
	}{
		"correct": {
			percentage: 50,
			duration:   time.Minute,
			tr:         isRaw(t, testResultOkStr, `{"id":123,"method":"bg_adjust_color","params":[50,60000]}`),
		},
		"wrong_percentage": {
			percentage: -150,
			duration:   time.Minute,
			expErr:     ErrWrongAdjustPercentage,
		},
		"wrong_duration": {
			percentage: 50,
			duration:   time.Millisecond,
			expErr:     ErrDurationTooSmall,
		},
		"err_connection": {
			percentage: 50,
			duration:   time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.AdjustBackgroundColor(testCtx, tt.percentage, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}
