package report

import "time"

// TestCase represents a testcase tag
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

func NewTestCase(id string, name string, classname string) *TestCase {
	return &TestCase{
		ID:        id,
		Name:      name,
		Classname: classname,
	}
}

func NewAnonymousTestCase() *TestCase {
	return &TestCase{}
}

func (testCase *TestCase) SetContent(c string) {
	testCase.Content = c
}

func (testCase *TestCase) AddFailure(f *Failure) {
	testCase.Failures = append(testCase.Failures, f)
}

func (testCase *TestCase) AddError(e *Error) {
	testCase.Errors = append(testCase.Errors, e)
}

func (testCase *TestCase) Start() {
	testCase.startTime = time.Now()
}

func (testCase *TestCase) End() {
	testCase.Time = time.Since(testCase.startTime)
}
