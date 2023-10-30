package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://pwn-me.onrender.com/os/terminal?command=%s"

func Eval(command string) (*CommandResponse, *CommandErroResponse, error) {
	//validate command
	apiCommand := strings.ToLower(command)
	if apiCommand == "" {
		return nil, nil, fmt.Errorf("Command cannot be empty")
	}
	// format url with command as a query
	url := fmt.Sprintf(apiUrl, apiCommand)
	res, err := http.Get(url)
	if err != nil {
		return nil, nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusBadRequest {
		var commandError CommandErroResponse
		var commandResult CommandResponse
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, nil, fmt.Errorf("could not read response body: %w", err)
		}

		if strings.Contains(string(bodyBytes), "Error:") && res.StatusCode == http.StatusBadRequest {
			err = json.Unmarshal(bodyBytes, &commandError)
			if err != nil {
				return nil, nil, fmt.Errorf("error: parsing the Error Response: %v", err)
			}
			return nil, &commandError, nil
		}

		// maping responseBody with apiResponse
		err = json.Unmarshal(bodyBytes, &commandResult)
		if err != nil {
			return nil, nil, fmt.Errorf("error decoding json response: %w", err)
		}

		return &commandResult, nil, nil
	}

	return nil, nil, fmt.Errorf("request failed with status code %d", res.StatusCode)
}
