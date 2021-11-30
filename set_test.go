package yeelight

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_SetColorTemperature(t *testing.T) {
	tests := map[string]struct {
		value    uint
		affect   string
		duration time.Duration
		tr       transportFn

		expErr error
	}{
		"correct": {
			value:    10,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"set_ct_abx","params":[10,"smooth",60000]}`, raw)

				return testResultOk, nil
			},
			expErr: nil,
		},
		"wrong_affect": {
			value:    10,
			affect:   "im_wrong_affect",
			duration: time.Minute,
			expErr:   ErrWrongAffect,
		},
		"wrong_duration": {
			value:    0,
			affect:   AffectSmooth,
			duration: 10,
			expErr:   ErrDurationTooSmall,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{transport: tt.tr}.SetColorTemperature(testCtx, testHost, testRequestID, tt.value, tt.affect, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetBackgroundColorTemperature(t *testing.T) {
	tests := map[string]struct {
		value    uint
		affect   string
		duration time.Duration
		tr       transportFn

		expErr error
	}{
		"correct": {
			value:    10,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"bg_set_ct_abx","params":[10,"smooth",60000]}`, raw)

				return testResultOk, nil
			},
			expErr: nil,
		},
		"wrong_affect": {
			value:    10,
			affect:   "im_wrong_affect",
			duration: time.Minute,
			expErr:   ErrWrongAffect,
		},
		"wrong_duration": {
			value:    0,
			affect:   AffectSmooth,
			duration: 10,
			expErr:   ErrDurationTooSmall,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			err := Client{transport: tt.tr}.SetBackgroundColorTemperature(testCtx, testHost, testRequestID, tt.value, tt.affect, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}
