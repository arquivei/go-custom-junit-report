package report

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestCase(t *testing.T) {
	actual := NewTestCase("id", "name", "classname")

	expected := &TestCase{
		ID:        "id",
		Name:      "name",
		Classname: "classname",
	}

	assert.Equal(t, expected, actual)
}

func TestNewAnonymousTestCase(t *testing.T) {
	actual := NewAnonymousTestCase()

	expected := &TestCase{}

	assert.Equal(t, expected, actual)
}

func TestSetContent(t *testing.T) {
	actual := NewAnonymousTestCase()
	actual.SetContent("content")

	expected := &TestCase{
		Content: "content",
	}

	assert.Equal(t, expected, actual)
}

func TestAddFailure_One(t *testing.T) {
	actual := NewAnonymousTestCase()
	actual.AddFailure(NewAnonymousFailure("content"))

	expected := &TestCase{
		Failures: []*Failure{
			{
				Content: "content",
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestAddFailure_Many(t *testing.T) {
	actual := NewAnonymousTestCase()
	actual.AddFailure(NewAnonymousFailure("content1"))
	actual.AddFailure(NewAnonymousFailure("content2"))
	actual.AddFailure(NewFailure("msg", "type", "content3"))

	expected := &TestCase{
		Failures: []*Failure{
			{
				Content: "content1",
			},
			{
				Content: "content2",
			},
			{
				Message: "msg",
				Type:    "type",
				Content: "content3",
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestAddError_One(t *testing.T) {
	actual := NewAnonymousTestCase()
	actual.AddError(NewAnonymousError("content"))

	expected := &TestCase{
		Errors: []*Error{
			{
				Content: "content",
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestAddError_Many(t *testing.T) {
	actual := NewAnonymousTestCase()
	actual.AddError(NewAnonymousError("content1"))
	actual.AddError(NewAnonymousError("content2"))
	actual.AddError(NewError("message", "type", "content3"))

	expected := &TestCase{
		Errors: []*Error{
			{
				Content: "content1",
			},
			{
				Content: "content2",
			},
			{
				Message: "message",
				Type:    "type",
				Content: "content3",
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestStart(t *testing.T) {
	actual := NewAnonymousTestCase()
	actual.Start()

	assert.NotNil(t, actual.startTime)
}

func TestEnd(t *testing.T) {
	actual := NewAnonymousTestCase()
	actual.Start()
	actual.End()

	assert.NotNil(t, actual.Time)
}
