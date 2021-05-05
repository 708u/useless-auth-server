package gateway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const resourceAPI = "/v1/resource/user"

type ResourceGateway struct {
	ResourceSrvURL string
}

type TmpResource struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *ResourceGateway) FetchUserResource(accessToken string) (string, error) {
	// TODO: hard code change path from arg
	req, _ := http.NewRequest(http.MethodPost, r.ResourceSrvURL, nil)
	req.URL.Path = resourceAPI
	req.Header.Set("Authorization", "Bearer "+accessToken)
	// api exec
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed ResourceGateway.FetchUserResource: %w", err)
	}

	if resp.StatusCode >= http.StatusMultipleChoices {
		return "", fmt.Errorf("failed ResourceGateway.FetchUserResource: v1/user returned response code %s", resp.Status)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// TODO: temporal body
	var tmp TmpResource
	if err := json.Unmarshal(respBody, &tmp); err != nil {
		return "", fmt.Errorf("failed ResourceGateway.FetchUserResource: %w", err)
	}
	return tmp.Name, nil
}
