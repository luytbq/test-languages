package types

type RunRequest struct {
	Submission string `json:"submission"`
}

type RunResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}
