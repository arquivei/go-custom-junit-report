package report

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

// TestSuites represents a testsuites tag
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

func NewTestSuites(id string, name string) *TestSuites {
	return &TestSuites{
		ID:       id,
		Name:     name,
		Tests:    0,
		Failures: 0,
		Errors:   0,
	}
}

func NewAnonymousTestSuites() *TestSuites {
	return &TestSuites{
		Tests:    0,
		Failures: 0,
		Errors:   0,
	}
}

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

func (suites *TestSuites) MakeReport() ([]byte, error) {
	suites.resolve()
	content, err := xml.MarshalIndent(suites, "", "    ")

	if err != nil {
		return []byte{}, err
	}

	return []byte(xml.Header + string(content)), nil
}

func (suites *TestSuites) SaveReport(filename string) error {
	content, err := suites.MakeReport()
	if err != nil {
		return err
	}

	return os.WriteFile(filename, content, 0644)
}

func (suites *TestSuites) RemoveTestSuite(id string) {
	for i, suite := range suites.TestSuites {
		if suite.ID == id {
			suites.TestSuites = append(suites.TestSuites[:i], suites.TestSuites[i+1:]...)
			return
		}
	}
}

func (suites *TestSuites) resolve() {
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
