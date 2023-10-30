package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://pwn-me.onrender.com/os/terminal?command=%s"

func Eval(command string) (*ApiResponse, error) {
	//validate command
	apiCommand := strings.ToLower(command)
	if apiCommand == "" {
		return nil, fmt.Errorf("Command cannot be empty")
	}
	// format url with command as a query
	url := fmt.Sprintf(apiUrl, apiCommand)
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()
	var commandResult ApiResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("could not read response body: %w", err)
		}
		// maping responseBody with apiResponse
		err = json.Unmarshal(bodyBytes, &commandResult)
		if err != nil {
			return nil, fmt.Errorf("error decoding json response: %w", err)
		}
	} else {
		return nil, fmt.Errorf("request failed with status code %d", res.StatusCode)
	}
	return &commandResult, nil
}
