package report

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	actual := NewError("message", "type", "content")

	expected := &Error{
		Message: "message",
		Type:    "type",
		Content: "content",
	}

	assert.Equal(t, expected, actual)
}

func TestNewAnonymousError(t *testing.T) {
	actual := NewAnonymousError("content")

	expected := &Error{
		Content: "content",
	}

	assert.Equal(t, expected, actual)
}
