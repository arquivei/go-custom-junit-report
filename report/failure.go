package report

type Failure struct {
	Message string `xml:"message,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Content string `xml:",chardata"`
}

func NewFailure(msg string, failureType string, content string) *Failure {
	return &Failure{
		Message: msg,
		Type:    failureType,
		Content: content,
	}
}

func NewAnonymousFailure(content string) *Failure {
	return &Failure{
		Content: content,
	}
}
