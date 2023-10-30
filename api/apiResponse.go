package api

type CommandResponse struct {
	CommandResult []string `json:"command-result"`
}

type CommandErroResponse struct {
	Error []string `json:"Error"`
}
