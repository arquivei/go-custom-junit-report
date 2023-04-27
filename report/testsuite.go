package report

import (
	"fmt"
	"time"
)

// TestSuite represents a testsuite tag
type TestSuite struct {
	ID        string        `xml:"id,attr,omitempty"`
	Name      string        `xml:"name,attr,omitempty"`
	Time      time.Duration `xml:"time,attr,omitempty"`
	Tests     int           `xml:"tests,attr"`
	Failures  int           `xml:"failures,attr"`
	Errors    int           `xml:"errors,attr"`
	TestCases []*TestCase   `xml:"testcase,omitempty"`
}

func NewTestSuite(id string, name string) *TestSuite {
	return &TestSuite{
		ID:       id,
		Name:     name,
		Tests:    0,
		Failures: 0,
		Errors:   0,
	}
}

func NewAnonymousTestSuite() *TestSuite {
	return &TestSuite{
		Tests:    0,
		Failures: 0,
		Errors:   0,
	}
}

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
