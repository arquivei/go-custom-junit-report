package report

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestSuites(t *testing.T) {
	actual := NewTestSuites("id", "name")

	expected := &TestSuites{
		ID:       "id",
		Name:     "name",
		Tests:    0,
		Failures: 0,
		Errors:   0,
	}

	assert.Equal(t, expected, actual)
}

func TestNewAnonymousTestSuites(t *testing.T) {
	actual := NewAnonymousTestSuites()

	expected := &TestSuites{
		Tests:    0,
		Failures: 0,
		Errors:   0,
	}

	assert.Equal(t, expected, actual)
}

func TestAddTestSuite_Success(t *testing.T) {
	actual := NewAnonymousTestSuites()
	err := actual.AddTestSuite(NewAnonymousTestSuite())
	assert.Nil(t, err)

	err = actual.AddTestSuite(NewAnonymousTestSuite())
	assert.Nil(t, err)

	err = actual.AddTestSuite(NewTestSuite("id", "name"))
	assert.Nil(t, err)

	assert.Equal(t, 3, len(actual.TestSuites))
}

func TestAddTestSuite_Error(t *testing.T) {
	actual := NewAnonymousTestSuites()
	err := actual.AddTestSuite(NewAnonymousTestSuite())
	assert.Nil(t, err)

	err = actual.AddTestSuite(NewTestSuite("id", "name"))
	assert.Nil(t, err)

	err = actual.AddTestSuite(NewTestSuite("id", "name"))
	assert.NotNil(t, err)

	assert.Equal(t, 2, len(actual.TestSuites))
}

func TestRemoveTestSuite(t *testing.T) {
	actual := NewAnonymousTestSuites()
	err := actual.AddTestSuite(NewTestSuite("id1", "name"))
	assert.Nil(t, err)

	err = actual.AddTestSuite(NewTestSuite("id2", "name"))
	assert.Nil(t, err)

	actual.RemoveTestSuite("id1")

	assert.Equal(t, 1, len(actual.TestSuites))
}

func TestMakeReport(t *testing.T) {
	suites := NewTestSuites("testsuites#1", "test_report")

	suite1 := NewTestSuite("testsuite#1", "suite 1")

	case1 := NewTestCase("case#1", "case 1", "report.TestMakeReport")

	err := suite1.AddTestCase(case1)
	assert.Nil(t, err)

	case2 := NewTestCase("case#2", "case 2", "report.TestMakeReport")

	f := NewFailure("msg1", "type_fail", "test failure 1")
	case2.AddFailure(f)
	f = NewFailure("msg2", "type_fail", "test failure 2")
	case2.AddFailure(f)
	f = NewFailure("msg3", "type_fail", "test failure 3")
	case2.AddFailure(f)

	e := NewError("msg4", "type_err", "test error 1")
	case2.AddError(e)
	e = NewError("msg5", "type_err", "test error 2")
	case2.AddError(e)
	e = NewError("msg6", "type_err", "test error 3")
	case2.AddError(e)

	err = suite1.AddTestCase(case2)
	assert.Nil(t, err)
	err = suites.AddTestSuite(suite1)
	assert.Nil(t, err)

	suite2 := NewAnonymousTestSuite()
	err = suite2.AddTestCase(NewAnonymousTestCase())
	assert.Nil(t, err)
	err = suite2.AddTestCase(NewAnonymousTestCase())
	assert.Nil(t, err)
	err = suites.AddTestSuite(suite2)
	assert.Nil(t, err)

	actual, err := suites.MakeReport()
	expected := mustLoadFile("make_report_expected.xml")

	assert.Nil(t, err)
	assert.Equal(t, string(expected), string(actual))
}

func TestResolve(t *testing.T) {
	suites := TestSuites{
		Tests:    0,
		Failures: 0,
		Errors:   0,
		Time:     0,
		TestSuites: []*TestSuite{
			{
				Time:     0,
				Tests:    0,
				Failures: 0,
				Errors:   0,
				TestCases: []*TestCase{
					{
						Time: 2,
						Failures: []*Failure{
							{
								Content: "c",
							},
						},
						Errors: []*Error{
							{
								Content: "c",
							},
						},
					},
					{
						Time: 12,
						Failures: []*Failure{
							{
								Content: "c",
							},
						},
					},
					{
						Time: 22,
						Errors: []*Error{
							{
								Content: "c",
							},
						},
					},
					{
						Time: 33,
					},
				},
			},
			{
				TestCases: []*TestCase{
					{
						Time: 31,
					},
				},
			},
		},
	}

	resolved := TestSuites{
		Tests:    5,
		Failures: 2,
		Errors:   2,
		Time:     100,
		TestSuites: []*TestSuite{
			{
				Time:     69,
				Tests:    4,
				Failures: 2,
				Errors:   2,
				TestCases: []*TestCase{
					{
						Time: 2,
						Failures: []*Failure{
							{
								Content: "c",
							},
						},
						Errors: []*Error{
							{
								Content: "c",
							},
						},
					},
					{
						Time: 12,
						Failures: []*Failure{
							{
								Content: "c",
							},
						},
					},
					{
						Time: 22,
						Errors: []*Error{
							{
								Content: "c",
							},
						},
					},
					{
						Time: 33,
					},
				},
			},
			{
				Time:     31,
				Tests:    1,
				Failures: 0,
				Errors:   0,
				TestCases: []*TestCase{
					{
						Time: 31,
					},
				},
			},
		},
	}

	suites.resolve()
	assert.Equal(t, resolved, suites)
}

func mustLoadFile(name string) []byte {
	b, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		panic(err)
	}

	return b
}
