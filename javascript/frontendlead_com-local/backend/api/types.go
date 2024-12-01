package api

type RunRequest struct {
	Slug       string `json:"slug"`
	Submission string `json:"submission"`
}

type RunResponse struct {
	State   string `json:"state"`
	Message string `json:"message,omitempty"`
}
