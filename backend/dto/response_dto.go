package dto

type CodeResponse struct {
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}