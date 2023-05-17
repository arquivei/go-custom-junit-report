package report

// Error corresponds to an error tag inside testcase and should be added every
// time an error happens during testing. A test case can have several errors.
// It has three fields: Message, Type, and Content. Message maps to the message
// optional attribute, which can take the error message. Type maps to the type
// optional attribute, which can take the error type, and Content maps to the
// tag's content text, which can be a detailed representation of the error, eg:
// message and stack trace.
type Error struct {
	Message string `xml:"message,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Content string `xml:",chardata"`
}

// NewError returns an Error with the given message, type, and content
func NewError(msg string, errorType string, content string) *Error {
	return &Error{
		Message: msg,
		Type:    errorType,
		Content: content,
	}
}

// NewAnonymousError returns an Error with the given content, but no message and
// type
func NewAnonymousError(content string) *Error {
	return &Error{
		Content: content,
	}
}
