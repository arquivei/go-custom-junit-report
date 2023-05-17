package report

import (
	"fmt"
	"time"
)

// TestSuite maps to a testsuite tag which represents a set of test cases. It
// has the following fields:
// ID:optional test suite ID. Maps to the id attribute. The attribute is omitted
// if empty. If given, an ID must be unique in the test suites.
// Name: optional test suite name. Maps to the name attribute. The attribute is
// omitted if empty.
// Time: optional duration of the suite. Maps to the time attribute. Omitted if
// empty. This field is calculated automatically by Testsuites.MakeReport().
// Tests: total amount of test cases in the suite. Maps to the tests attribute.
// This field is calculated automatically by Testsuites.MakeReport().
// Failures: total amount of failures cases in the suite. Maps to the failures
// attribute. This field is calculated automatically by Testsuites.MakeReport().
// Errors: total amount of errors in the suite. Maps to the errors attribute.
// This field is calculated automatically by Testsuites.MakeReport().
// TestCases: test cases in the suite. Each element maps to its own testcase
// tag.
type TestSuite struct {
	ID        string        `xml:"id,attr,omitempty"`
	Name      string        `xml:"name,attr,omitempty"`
	Time      time.Duration `xml:"time,attr,omitempty"`
	Tests     int           `xml:"tests,attr"`
	Failures  int           `xml:"failures,attr"`
	Errors    int           `xml:"errors,attr"`
	TestCases []*TestCase   `xml:"testcase,omitempty"`
}

// NewTestSuite returns a new TestSuite with the given id and name
func NewTestSuite(id string, name string) *TestSuite {
	return &TestSuite{
		ID:       id,
		Name:     name,
		Tests:    0,
		Failures: 0,
		Errors:   0,
	}
}

// NewAnonymousTestSuite returns a new empty TestSuite
func NewAnonymousTestSuite() *TestSuite {
	return &TestSuite{
		Tests:    0,
		Failures: 0,
		Errors:   0,
	}
}

// AddTestCase adds a TestCase to the suite. If the test case has an ID, it must
// be unique within the suite. If it isn't an error is returned.
func (suite *TestSuite) AddTestCase(testcase *TestCase) error {
	if len(testcase.ID) > 0 {
		for _, c := range suite.TestCases {
			if c.ID == testcase.ID {
				return fmt.Errorf(
					"cannot add test case: suite ID=%s already contains a case with ID=%s",
					suite.ID,
					testcase.ID,
				)
			}
		}
	}

	suite.TestCases = append(suite.TestCases, testcase)
	return nil
}

// RemoveTestCase removes a test case with the given id from the suite if it
// exists
func (suite *TestSuite) RemoveTestCase(id string) {
	for i, testcase := range suite.TestCases {
		if testcase.ID == id {
			suite.TestCases = append(
				suite.TestCases[:i],
				suite.TestCases[i+1:]...,
			)
			return
		}
	}
}
