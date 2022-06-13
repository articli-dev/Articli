package utils

type StatusCode int

func (c StatusCode) IsSuccess() bool {
	return c >= 200 && c < 300 || c == 304
}
