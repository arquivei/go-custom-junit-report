package report

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

// TestSuites maps to a testsuites tag which represents a set of test suites. It
// has the following fields:
// ID: optional ID. Maps to the id attribute. The attribute is omitted if empty.
// Name: optional name. Maps to the name attribute. The attribute is omitted if
// empty.
// Tests: total amount of test cases. Maps to the tests attribute. This field is
// calculated automatically by Testsuites.MakeReport().
// Failures: total amount of failures. Maps to the failures attribute. This
// field is calculated automatically by Testsuites.MakeReport().
// Errors: total amount of errors. Maps to the errors attribute. This field is
// calculated automatically by Testsuites.MakeReport().
// Time: optional duration of the test. Maps to the time attribute. Omitted if
// empty.
// TestSuites: test suites. Each element maps to its own testsuite tag.
type TestSuites struct {
	XMLName    xml.Name      `xml:"testsuites"`
	ID         string        `xml:"id,attr,omitempty"`
	Name       string        `xml:"name,attr,omitempty"`
	Tests      int           `xml:"tests,attr"`
	Failures   int           `xml:"failures,attr"`
	Errors     int           `xml:"errors,attr"`
	Time       time.Duration `xml:"time,attr,omitempty"`
	TestSuites []*TestSuite  `xml:"testsuite,omitempty"`
}

// NewTestSuites creates a new TestSuites with the given id and name
func NewTestSuites(id string, name string) *TestSuites {
	return &TestSuites{
		ID:       id,
		Name:     name,
		Tests:    0,
		Failures: 0,
		Errors:   0,
	}
}

// NewAnonymousTestSuites returns an empty TestSuites
func NewAnonymousTestSuites() *TestSuites {
	return &TestSuites{
		Tests:    0,
		Failures: 0,
		Errors:   0,
	}
}

// AddTestSuite add a TestSuite to the TestSuites. If the test suite has an ID,
// it must be unique within the suites. If it isn't an error is returned.
func (suites *TestSuites) AddTestSuite(suite *TestSuite) error {
	if len(suite.ID) > 0 {
		for _, s := range suites.TestSuites {
			if s.ID == suite.ID {
				return fmt.Errorf(
					"cannot add test suite: suites ID=%s already contains a suite with ID=%s",
					suites.ID,
					suite.ID,
				)
			}
		}
	}

	suites.TestSuites = append(suites.TestSuites, suite)
	return nil
}

// MakeReport generates the report XML as a slice of bytes. It is useful for any
// output other than generating a file. For saving the report as a file you
// should use SaveReport instead. All values are automatically calculated when
// calling this method.
func (suites *TestSuites) MakeReport() ([]byte, error) {
	suites.resolve()
	content, err := xml.MarshalIndent(suites, "", "    ")

	if err != nil {
		return []byte{}, err
	}

	return []byte(xml.Header + string(content)), nil
}

// SaveReport saves the report XMl in the given file name with the 644
// permission settings. All values are automatically calculated when calling
// this method.
func (suites *TestSuites) SaveReport(filename string) error {
	content, err := suites.MakeReport()
	if err != nil {
		return err
	}

	return os.WriteFile(filename, content, 0644)
}

// RemoveTestSuite removes a suite with the given id if it exists.
func (suites *TestSuites) RemoveTestSuite(id string) {
	for i, suite := range suites.TestSuites {
		if suite.ID == id {
			suites.TestSuites = append(suites.TestSuites[:i], suites.TestSuites[i+1:]...)
			return
		}
	}
}

// resolve calculates all the automatically calculated values (total time,
// amount of tests, errors, etc.). This method resets the values before each
// calculation so it is safe to call it multiple times.
func (suites *TestSuites) resolve() {
	suites.reset()
	for _, suite := range suites.TestSuites {
		for _, testCase := range suite.TestCases {
			suite.Tests++
			suite.Failures += len(testCase.Failures)
			suite.Errors += len(testCase.Errors)
			suite.Time += testCase.Time
		}

		suites.Tests += suite.Tests
		suites.Failures += suite.Failures
		suites.Errors += suite.Errors
		suites.Time += suite.Time
	}
}

// reset sets all automatically calculated values to 0
func (suites *TestSuites) reset() {
	for _, suite := range suites.TestSuites {
		suite.Tests = 0
		suite.Failures = 0
		suite.Errors = 0
		suite.Time = 0
	}

	suites.Tests = 0
	suites.Failures = 0
	suites.Errors = 0
	suites.Time = 0
}
