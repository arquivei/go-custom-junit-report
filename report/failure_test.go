package report

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFailure(t *testing.T) {
	actual := NewFailure("message", "type", "content")

	expected := &Failure{
		Message: "message",
		Type:    "type",
		Content: "content",
	}

	assert.Equal(t, expected, actual)
}

func TestNewAnonymousFailure(t *testing.T) {
	actual := NewAnonymousFailure("content")

	expected := &Failure{
		Content: "content",
	}

	assert.Equal(t, expected, actual)
}
