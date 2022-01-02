package yeelight

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testCtx         = context.Background()
	testHost        = "im_test_home"
	testResultOk    = []byte(`{"id":123, "result":["ok"]}`) // deprecated
	testResultOkStr = `{"id":123, "result":["ok"]}`
)

func TestClient_Get(t *testing.T) {
	tests := map[string]struct {
		properties []string
		tr         transportFn

		exp    map[string]string
		expErr error
	}{
		"correct": {
			properties: []string{PropertyPower, PropertySat},
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				assert.Equal(t, testCtx, ctx)
				assert.Equal(t, testHost, host)
				assert.Equal(t, `{"id":123,"method":"get_prop","params":["power","sat"]}`, raw)

				return []byte(`{"id":1, "result":["on", "100"]}`), nil
			},
			exp: map[string]string{PropertyPower: "on", PropertySat: "100"},
		},
		"wrong_number_of_result_items": {
			properties: []string{PropertyPower, PropertySat},
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return []byte(`{"id":1, "result":["on"]}`), nil
			},
			expErr: ErrWrongNumberOfResultItems,
		},
		"empty_properties_list": {
			properties: []string{},
			exp:        map[string]string{},
		},
		"err_connection": {
			properties: []string{PropertyPower, PropertySat},
			tr: func(ctx context.Context, host string, raw string) ([]byte, error) {
				return nil, ErrConnect
			},
			expErr: ErrConnect,
		},
	}

	for testCase, tt := range tests {
		t.Run(testCase, func(t *testing.T) {
			got, err := Client{host: testHost, transport: tt.tr}.GetProperties(testCtx, tt.properties)

			require.Equal(t, tt.expErr, err)
			require.Equal(t, tt.exp, got)
		})
	}
}
