package report

// Failure corresponds to a failure tag inside testcase and should be added
// every time a failure happens during testing. A test case can have several
// failures.
// It has three fields: Message, Type, and Content. Message maps to the message
// optional attribute, which can take the failure message. Type maps to the type
// optional attribute, which can take the failure type, and Content maps to the
// tag's content text, which can be a detailed representation of the failure,
// eg: message, class, and file.
type Failure struct {
	Message string `xml:"message,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Content string `xml:",chardata"`
}

// NewFailure returns a Failure with the given message, type, and content
func NewFailure(msg string, failureType string, content string) *Failure {
	return &Failure{
		Message: msg,
		Type:    failureType,
		Content: content,
	}
}

// NewAnonymousError returns a Failure with the given content, but no message
// and type
func NewAnonymousFailure(content string) *Failure {
	return &Failure{
		Content: content,
	}
}
