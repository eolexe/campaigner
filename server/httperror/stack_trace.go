package httperror

type StackTraceElement struct {
	LineNumber int    `json:"line_number"`
	ClassName  string `json:"class_name"`
	FileName   string `json:"file_name"`
	MethodName string `json:"method_name"`
}

type StackTrace []StackTraceElement

func (s *StackTrace) AddEntry(lineNumber int, packageName, fileName, methodName string) {
	*s = append(*s, StackTraceElement{lineNumber, packageName, fileName, methodName})
}
