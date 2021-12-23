package yeelight

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func isRaw(t *testing.T, resp string, exp string) transportFn {
	return func(ctx context.Context, host string, raw string) ([]byte, error) {
		assert.Equal(t, testCtx, ctx)
		assert.Equal(t, testHost, host)
		assert.Equal(t, exp, raw)

		return []byte(resp), nil
	}
}
