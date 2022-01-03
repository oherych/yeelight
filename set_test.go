package yeelight

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_SetName(t *testing.T) {
	tests := map[string]struct {
		name string
		tr   transportFn

		expErr error
	}{
		"correct": {
			name: "im_new_name",
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"set_name","params":["im_new_name"]}`, raw)

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
			err := Client{host: testHost, transport: tt.tr}.SetName(testCtx, tt.name)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetColorTemperature(t *testing.T) {
	tests := map[string]struct {
		value    int
		affect   string
		duration time.Duration
		tr       transportFn

		expErr error
	}{
		"correct": {
			value:    2000,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"set_ct_abx","params":[2000,"smooth",60000]}`, raw)

				return testResultOk, nil
			},
			expErr: nil,
		},
		"wrong_affect": {
			value:    2000,
			affect:   "im_wrong_affect",
			duration: time.Minute,
			expErr:   ErrWrongAffect,
		},
		"wrong_duration": {
			value:    2000,
			affect:   AffectSmooth,
			duration: 10,
			expErr:   ErrDurationTooSmall,
		},
		"wrong_color_temperature": {
			value:    1500,
			affect:   AffectSmooth,
			duration: time.Minute,
			expErr:   ErrWrongColorTemperature,
		},
		"err_connection": {
			value:    2000,
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
			err := Client{host: testHost, transport: tt.tr}.SetColorTemperature(testCtx, tt.value, tt.affect, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetBackgroundColorTemperature(t *testing.T) {
	tests := map[string]struct {
		value    int
		affect   string
		duration time.Duration
		tr       transportFn

		expErr error
	}{
		"correct": {
			value:    2000,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"bg_set_ct_abx","params":[2000,"smooth",60000]}`, raw)

				return testResultOk, nil
			},
			expErr: nil,
		},
		"wrong_affect": {
			value:    2000,
			affect:   "im_wrong_affect",
			duration: time.Minute,
			expErr:   ErrWrongAffect,
		},
		"wrong_duration": {
			value:    2000,
			affect:   AffectSmooth,
			duration: 10,
			expErr:   ErrDurationTooSmall,
		},
		"wrong_color_temperature": {
			value:    1500,
			affect:   AffectSmooth,
			duration: time.Minute,
			expErr:   ErrWrongColorTemperature,
		},
		"err_connection": {
			value:    2000,
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
			err := Client{host: testHost, transport: tt.tr}.SetBackgroundColorTemperature(testCtx, tt.value, tt.affect, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetRGB(t *testing.T) {
	tests := map[string]struct {
		value    int
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
				assert.Equal(t, `{"id":123,"method":"set_rgb","params":[10,"smooth",60000]}`, raw)

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
		"wrong_rgb": {
			value:    20000000,
			affect:   AffectSmooth,
			duration: time.Minute,
			expErr:   ErrWrongRGBValue,
		},
		"err_connection": {
			value:    10,
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
			err := Client{host: testHost, transport: tt.tr}.SetRGB(testCtx, tt.value, tt.affect, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetBackgroundRGB(t *testing.T) {
	tests := map[string]struct {
		value    int
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
				assert.Equal(t, `{"id":123,"method":"bg_set_rgb","params":[10,"smooth",60000]}`, raw)

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
		"wrong_rgb": {
			value:    20000000,
			affect:   AffectSmooth,
			duration: time.Minute,
			expErr:   ErrWrongRGBValue,
		},
		"err_connection": {
			value:    10,
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
			err := Client{host: testHost, transport: tt.tr}.SetBackgroundRGB(testCtx, tt.value, tt.affect, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetHSV(t *testing.T) {
	tests := map[string]struct {
		hue      int
		sat      int
		affect   string
		duration time.Duration
		tr       transportFn

		expErr error
	}{
		"correct": {
			hue:      10,
			sat:      20,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"set_hsv","params":[10,20,"smooth",60000]}`, raw)

				return testResultOk, nil
			},
			expErr: nil,
		},
		"wrong_affect": {
			hue:      10,
			sat:      20,
			affect:   "im_wrong_affect",
			duration: time.Minute,
			expErr:   ErrWrongAffect,
		},
		"wrong_duration": {
			hue:      10,
			sat:      20,
			affect:   AffectSmooth,
			duration: 10,
			expErr:   ErrDurationTooSmall,
		},
		"wrong_hue": {
			hue:      9999,
			sat:      20,
			affect:   AffectSmooth,
			duration: time.Minute,
			expErr:   ErrWrongHueValue,
		},
		"wrong_sat": {
			hue:      10,
			sat:      9999,
			affect:   AffectSmooth,
			duration: time.Minute,
			expErr:   ErrWrongSatValue,
		},
		"err_connection": {
			hue:      10,
			sat:      20,
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
			err := Client{host: testHost, transport: tt.tr}.SetHSV(testCtx, tt.hue, tt.sat, tt.affect, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetBackgroundHSV(t *testing.T) {
	tests := map[string]struct {
		hue      int
		sat      int
		affect   string
		duration time.Duration
		tr       transportFn

		expErr error
	}{
		"correct": {
			hue:      10,
			sat:      20,
			affect:   AffectSmooth,
			duration: time.Minute,
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"bg_set_hsv","params":[10,20,"smooth",60000]}`, raw)

				return testResultOk, nil
			},
			expErr: nil,
		},
		"wrong_affect": {
			hue:      10,
			sat:      20,
			affect:   "im_wrong_affect",
			duration: time.Minute,
			expErr:   ErrWrongAffect,
		},
		"wrong_duration": {
			hue:      10,
			sat:      20,
			affect:   AffectSmooth,
			duration: 10,
			expErr:   ErrDurationTooSmall,
		},
		"wrong_hue": {
			hue:      9999,
			sat:      20,
			affect:   AffectSmooth,
			duration: time.Minute,
			expErr:   ErrWrongHueValue,
		},
		"wrong_sat": {
			hue:      10,
			sat:      9999,
			affect:   AffectSmooth,
			duration: time.Minute,
			expErr:   ErrWrongSatValue,
		},
		"err_connection": {
			hue:      10,
			sat:      20,
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
			err := Client{host: testHost, transport: tt.tr}.SetBackgroundHSV(testCtx, tt.hue, tt.sat, tt.affect, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetBright(t *testing.T) {
	tests := map[string]struct {
		value    int
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
				assert.Equal(t, `{"id":123,"method":"set_bright","params":[10,"smooth",60000]}`, raw)

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
			value:    10,
			affect:   AffectSmooth,
			duration: 10,
			expErr:   ErrDurationTooSmall,
		},
		"wrong_value": {
			value:    101,
			affect:   AffectSmooth,
			duration: time.Minute,
			expErr:   ErrWrongBrightValue,
		},
		"err_connection": {
			value:    10,
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
			err := Client{host: testHost, transport: tt.tr}.SetBright(testCtx, tt.value, tt.affect, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetBackgroundBright(t *testing.T) {
	tests := map[string]struct {
		value    int
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
				assert.Equal(t, `{"id":123,"method":"bg_set_bright","params":[10,"smooth",60000]}`, raw)

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
			value:    10,
			affect:   AffectSmooth,
			duration: 10,
			expErr:   ErrDurationTooSmall,
		},
		"wrong_value": {
			value:    101,
			affect:   AffectSmooth,
			duration: time.Minute,
			expErr:   ErrWrongBrightValue,
		},
		"err_connection": {
			value:    10,
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
			err := Client{host: testHost, transport: tt.tr}.SetBackgroundBright(testCtx, tt.value, tt.affect, tt.duration)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetDefault(t *testing.T) {
	tests := map[string]struct {
		tr transportFn

		expErr error
	}{
		"correct": {
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"set_default","params":[]}`, raw)

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
			err := Client{host: testHost, transport: tt.tr}.SetDefault(testCtx)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetBackgroundDefault(t *testing.T) {
	tests := map[string]struct {
		tr transportFn

		expErr error
	}{
		"correct": {
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"bg_set_default","params":[]}`, raw)

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
			err := Client{host: testHost, transport: tt.tr}.SetBackgroundDefault(testCtx)

			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestClient_SetMusic(t *testing.T) {
	const testMusicHost = "test_music_host"
	const testMusicPost = 1234

	tests := map[string]struct {
		on bool

		tr transportFn

		expErr error
	}{
		"correct_on": {
			on: true,
			tr: isRaw(t, testResultOkStr, `{"id":123,"method":"set_music","params":[1,"test_music_host",1234]}`),
		},
		"correct_off": {
			on: false,
			tr: isRaw(t, testResultOkStr, `{"id":123,"method":"set_music","params":[0]}`),
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
			err := Client{host: testHost, transport: tt.tr}.SetMusic(testCtx, tt.on, testMusicHost, testMusicPost)

			require.Equal(t, tt.expErr, err)
		})
	}
}
