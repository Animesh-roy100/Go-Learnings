package dto

// CommitChangesRequest represents the input for committing changes
type CommitChangesRequest struct {
	Changes       map[string]interface{} `json:"changes"`
	CommitMessage string                 `json:"commit_message,omitempty"`
}
