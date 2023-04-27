package report

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestSuite(t *testing.T) {
	actual := NewTestSuite("id", "name")

	expected := &TestSuite{
		ID:       "id",
		Name:     "name",
		Tests:    0,
		Failures: 0,
		Errors:   0,
	}

	assert.Equal(t, expected, actual)
}

func TestNewAnonymousTestSuite(t *testing.T) {
	actual := NewAnonymousTestSuite()

	expected := &TestSuite{
		Tests:    0,
		Failures: 0,
		Errors:   0,
	}

	assert.Equal(t, expected, actual)
}

func TestAddTestCase_Success(t *testing.T) {
	actual := NewAnonymousTestSuite()
	err := actual.AddTestCase(NewAnonymousTestCase())
	assert.Nil(t, err)

	err = actual.AddTestCase(NewAnonymousTestCase())
	assert.Nil(t, err)

	err = actual.AddTestCase(NewAnonymousTestCase())
	assert.Nil(t, err)

	assert.Equal(t, 3, len(actual.TestCases))
}

func TestAddTestCase_Error(t *testing.T) {
	actual := NewAnonymousTestSuite()
	err := actual.AddTestCase(NewTestCase("1", "name", "class"))
	assert.Nil(t, err)

	err = actual.AddTestCase(NewTestCase("1", "name", "class"))
	assert.NotNil(t, err)

	err = actual.AddTestCase(NewTestCase("2", "name", "class"))
	assert.Nil(t, err)

	assert.Equal(t, 2, len(actual.TestCases))
}

func TestRemoveTestCase(t *testing.T) {
	actual := NewAnonymousTestSuite()
	err := actual.AddTestCase(NewTestCase("1", "name", "class"))
	assert.Nil(t, err)

	err = actual.AddTestCase(NewTestCase("2", "name", "class"))
	assert.Nil(t, err)

	actual.RemoveTestCase("1")

	assert.Equal(t, 1, len(actual.TestCases))
}
