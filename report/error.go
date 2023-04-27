package report

type Error struct {
	Message string `xml:"message,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Content string `xml:",chardata"`
}

func NewError(msg string, errorType string, content string) *Error {
	return &Error{
		Message: msg,
		Type:    errorType,
		Content: content,
	}
}

func NewAnonymousError(content string) *Error {
	return &Error{
		Content: content,
	}
}
