package segmentfault

type APIError struct {
	StatusCode int    `json:"status_code"`
	Content    string `json:"content"`
}

func (e *APIError) Error() string {
	return e.Content
}
