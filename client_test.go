package yeelight

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	const testHost = "im_test_host"

	client := New(testHost)

	assert.Equal(t, testHost, client.host)
	assert.Equal(t, reflect.ValueOf(defaultTransport).Pointer(), reflect.ValueOf(client.transport).Pointer())
}

func isRaw(t *testing.T, resp string, exp string) transportFn {
	return func(ctx context.Context, host string, raw string) ([]byte, error) {
		assert.Equal(t, testCtx, ctx)
		assert.Equal(t, testHost, host)
		assert.Equal(t, exp, raw)

		return []byte(resp), nil
	}
}
