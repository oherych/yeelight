package yeelight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknownError_Error(t *testing.T) {
	assert.Equal(t, "im_test_error", UnknownError("im_test_error").Error())

}
