package report

import "time"

// TestCase maps to a testcase tag which represents a test case. It has
// the following fields:
// ID: optional test case ID. Maps to the id attribute. The attribute is omitted
// if empty. If given, an ID must be unique in the test suite.
// Name: optional test case name. Maps to the name attribute. The attribute is
// omitted if empty.
// Time: optional duration of the test. Maps to the time attribute. Omitted if
// empty.
// Classname: optional name of the module beiong tested. Maps to the classname
// attribute. Omitted if empty.
// Content: optional text content of the test. Maps to the content of the tag.
// Failures: test failures. Each element maps to its own failure tag.
// Errors: test errors. Each element maps to its own error tag.
type TestCase struct {
	ID        string        `xml:"id,attr,omitempty"`
	Name      string        `xml:"name,attr,omitempty"`
	Time      time.Duration `xml:"time,attr,omitempty"`
	Classname string        `xml:"classname,attr,omitempty"`
	Content   string        `xml:",chardata"`
	Failures  []*Failure    `xml:"failure"`
	Errors    []*Error      `xml:"error"`
	startTime time.Time     `xml:"-"`
}

// NewTestCase returns a test case with the given id, name, and classname
func NewTestCase(id string, name string, classname string) *TestCase {
	return &TestCase{
		ID:        id,
		Name:      name,
		Classname: classname,
	}
}

// NewAnonymousTestCase returns an empty test case
func NewAnonymousTestCase() *TestCase {
	return &TestCase{}
}

// SetContent sets the test case content
func (testCase *TestCase) SetContent(c string) {
	testCase.Content = c
}

// AddFailure adds a failure to the test case
func (testCase *TestCase) AddFailure(f *Failure) {
	testCase.Failures = append(testCase.Failures, f)
}

// AddError adds an error to the test case
func (testCase *TestCase) AddError(e *Error) {
	testCase.Errors = append(testCase.Errors, e)
}

// Start sets the time the test started
func (testCase *TestCase) Start() {
	testCase.startTime = time.Now()
}

// End sets the test cases duration
func (testCase *TestCase) End() {
	testCase.Time = time.Since(testCase.startTime)
}
