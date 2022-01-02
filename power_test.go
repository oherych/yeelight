package yeelight

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Power(t *testing.T) {
	tests := map[string]struct {
		mode     uint
		on       bool
		affect   string
		duration time.Duration

		tr transportFn

		expErr error
	}{
		"correct_on": {
			mode:     PowerModeDefault,
			on:       true,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"set_power","params":["on","smooth",60000,0]}`, raw)

				return testResultOk, nil
			},
			expErr: nil,
		},
		"correct_on_mode": {
			mode:     PowerModeRGB,
			on:       true,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"set_power","params":["on","smooth",60000,2]}`, raw)

				return testResultOk, nil
			},
			expErr: nil,
		},
		"correct_off": {
			mode:     PowerModeColorFlow,
			on:       false,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"set_power","params":["off","smooth",60000]}`, raw)

				return testResultOk, nil
			},
			expErr: nil,
		},
		"wrong_affect": {
			mode:     PowerModeDefault,
			on:       true,
			affect:   "im_wrong_affect",
			duration: time.Minute,
			expErr:   ErrWrongAffect,
		},
		"wrong_duration": {
			mode:     PowerModeDefault,
			on:       true,
			affect:   AffectSmooth,
			duration: 10,
			expErr:   ErrDurationTooSmall,
		},
		"err_connection": {
			mode:     PowerModeDefault,
			on:       true,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.Power(testCtx, tt.on, tt.mode, tt.affect, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_BackgroundPower(t *testing.T) {
	tests := map[string]struct {
		mode     uint
		on       bool
		affect   string
		duration time.Duration

		tr transportFn

		expErr error
	}{
		"correct_on": {
			mode:     PowerModeDefault,
			on:       true,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"bg_set_power","params":["on","smooth",60000,0]}`, raw)

				return testResultOk, nil
			},
			expErr: nil,
		},
		"correct_on_mode": {
			mode:     PowerModeRGB,
			on:       true,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"bg_set_power","params":["on","smooth",60000,2]}`, raw)

				return testResultOk, nil
			},
			expErr: nil,
		},
		"correct_off": {
			mode:     PowerModeColorFlow,
			on:       false,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"bg_set_power","params":["off","smooth",60000]}`, raw)

				return testResultOk, nil
			},
			expErr: nil,
		},
		"wrong_affect": {
			mode:     PowerModeDefault,
			on:       true,
			affect:   "im_wrong_affect",
			duration: time.Minute,
			expErr:   ErrWrongAffect,
		},
		"wrong_duration": {
			mode:     PowerModeDefault,
			on:       true,
			affect:   AffectSmooth,
			duration: 10,
			expErr:   ErrDurationTooSmall,
		},
		"err_connection": {
			mode:     PowerModeDefault,
			on:       true,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{host: testHost, transport: tt.tr}.BackgroundPower(testCtx, tt.on, tt.mode, tt.affect, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_Toggle(t *testing.T) {
	tests := map[string]struct {
		tr transportFn

		expErr error
	}{
		"correct": {
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"toggle","params":[]}`, raw)

				return testResultOk, nil
			},
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
			err := Client{host: testHost, transport: tt.tr}.Toggle(testCtx)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_BackgroundToggle(t *testing.T) {
	tests := map[string]struct {
		tr transportFn

		expErr error
	}{
		"correct": {
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"bg_toggle","params":[]}`, raw)

				return testResultOk, nil
			},
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
			err := Client{host: testHost, transport: tt.tr}.BackgroundToggle(testCtx)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_DevToggle(t *testing.T) {
	tests := map[string]struct {
		tr transportFn

		expErr error
	}{
		"correct": {
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"dev_toggle","params":[]}`, raw)

				return testResultOk, nil
			},
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
			err := Client{host: testHost, transport: tt.tr}.DevToggle(testCtx)

			require.Equal(t, tt.expErr, err)
		})
	}
}
